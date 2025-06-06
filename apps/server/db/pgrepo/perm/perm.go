package perm

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"slices"
	"strings"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"go.uber.org/zap"
)

//go:embed model.conf
var authModel string

//go:embed predefined_policy.csv
var predefinedPolicy string

func readPredefiendPolicy() (p [][]string, g [][]string, g2 [][]string) {
	lines := strings.Split(predefinedPolicy, "\n")
	slices.DeleteFunc(lines, func(line string) bool {
		return strings.TrimSpace(line) == ""
	})
	var policies [][]string
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		policy := strings.Split(l, ",")
		for i, v := range policy {
			policy[i] = strings.TrimSpace(v)
		}
		policies = append(policies, policy)
	}

	// var p, g, g2 [][]string
	for _, r := range policies {
		switch r[0] {
		case "p":
			p = append(p, r[1:])
		case "g":
			g = append(g, r[1:])
		case "g2":
			g2 = append(g2, r[1:])
		}
	}
	return p, g, g2

}

type Perm struct {
	enforcer *casbin.Enforcer
	log      *zap.Logger
}

func New(db *sql.DB, dbType string, log *zap.Logger) (*Perm, error) {
	log = log.Named("Perm")
	m, err := model.NewModelFromString(authModel)
	if err != nil {
		return nil, err
	}
	a, err := sqladapter.NewAdapter(db, dbType, "casbin_rule")
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}
	if err = e.LoadPolicy(); err != nil {
		return nil, err
	}
	return &Perm{enforcer: e, log: log}, nil
}

func InitPerm(perm *Perm, db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM casbin_rule").Scan(&count)
	if err != nil {
		perm.log.Fatal("failed to check rule table", zap.Error(err))
	}
	if count == 0 {
		perm.log.Info("rule table empty,filling predefined policy")
		// perm.log.Debug("check embled policy content", zap.String("predefinedPolicy", predefinedPolicy))
		p, g, g2 := readPredefiendPolicy()
		// perm.log.Debug("check embled policy content", zap.Any("policies", policies))

		result, err := perm.enforcer.AddNamedPolicies("p", p)
		if err != nil {
			return fmt.Errorf("failed to add polices:%w", err)
		}
		if !result {
			return errors.New("failed to add polices:rule already exists?")
		}

		result, err = perm.enforcer.AddNamedGroupingPolicies("g", g)
		if err != nil {
			return fmt.Errorf("failed to add polices:%w", err)
		}
		if !result {
			return errors.New("failed to add polices:rule already exists?")
		}

		result, err = perm.enforcer.AddNamedGroupingPolicies("g2", g2)
		if err != nil {
			return fmt.Errorf("failed to add polices:%w", err)
		}
		if !result {
			return errors.New("failed to add polices:rule already exists?")
		}

		// perm.enforcer.SavePolicy()
		perm.log.Info("policy table init complete")
	}
	return nil
}

func (p *Perm) Check(sub ResId, res ResId, act ActStr) (bool, error) {
	return p.enforcer.Enforce(sub.Str(), res.Type().Act(act), res.Str())
}

func (p *Perm) AddResRelation(parent ResId, child ResId) {
	p.enforcer.AddNamedGroupingPolicy("g2", child.Str(), parent.Str())
}

func (p *Perm) DeleteResRelation(res ResId) (bool, error) {
	result, err := p.enforcer.RemoveFilteredNamedGroupingPolicy("g2", 0, res.Str())
	if err != nil {
		return result, err
	}
	result2, err := p.enforcer.RemoveFilteredNamedGroupingPolicy("g2", 1, res.Str())
	return result || result2, err
}

func (p *Perm) AddPerm(sub ResId, role RoleStr, res ResId) {
	p.enforcer.AddPolicy(res.Str(), res.Type().Role(role), res.Str())
}
func (p *Perm) RemovePerm(sub ResId) {
	p.enforcer.RemoveFilteredPolicy(0, sub.Str())
}

func (p *Perm) Load() error {
	return p.enforcer.LoadPolicy()
}

func (p *Perm) Save() error {
	return p.enforcer.SavePolicy()
}
