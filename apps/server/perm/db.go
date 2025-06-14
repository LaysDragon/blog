package perm

import (
	"database/sql"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

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
	} else {
		perm.log.Info("no need init policy table")
	}
	return nil
}
