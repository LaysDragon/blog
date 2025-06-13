package perm

import (
	_ "embed"
	"fmt"
	"log"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	stringadapter "github.com/casbin/casbin/v2/persist/string-adapter"
	"github.com/casbin/casbin/v2/util"
	// scas "github.com/qiangmzsx/string-adapter"
)

// //go:embed model.conf
// var authModel string

// //go:embed predefined_policy.csv
// var predefinedPolicy string

//go:embed test_policy.csv
var testPolicy string

func initCasbin() *casbin.Enforcer {
	m, err := model.NewModelFromString(authModel)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)

	}
	sa := stringadapter.NewAdapter(predefinedPolicy + "\n" + testPolicy)
	e, err := casbin.NewEnforcer(m, sa)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch)
	e.AddNamedMatchingFunc("g2", "KeyMatch2", util.KeyMatch)
	return e
}

func TestRule(t *testing.T) {
	e := initCasbin()

	tests := []struct {
		sub  string
		act  string
		res  string
		want bool
	}{
		{
			sub:  "user.1",
			act:  "ACT::POST/WRITE",
			res:  "post.1",
			want: true,
		},
		{
			sub:  "user.1",
			act:  "ACT::POST/WRITE",
			res:  "post.2",
			want: true,
		},
		{
			sub:  "user.2",
			act:  "ACT::POST/WRITE",
			res:  "post.3",
			want: true,
		},
		{
			sub:  "user.2",
			act:  "ACT::POST/WRITE",
			res:  "post.1",
			want: false,
		},
		{
			sub:  "user.2",
			act:  "ACT::POST/READ",
			res:  "post.1",
			want: true,
		},
		{
			sub:  "user.admin",
			act:  "ACT::POST/WRITE",
			res:  "post.1",
			want: true,
		},
		{
			sub:  "user.anon",
			act:  "ACT::POST/READ",
			res:  "post.1",
			want: true,
		},
		{
			sub:  "user.anon",
			act:  "ACT::POST/WRITE",
			res:  "post.1",
			want: false,
		},
		{
			sub:  "user.anon",
			act:  "ACT::COMMENT/WRITE",
			res:  "post.1",
			want: true,
		},
		{
			sub:  "user.anon",
			act:  "ACT::COMMENT/READ",
			res:  "comment.1",
			want: true,
		},
		{
			sub:  "user.admin",
			act:  "ACT::USER_ADMIN/WRITE",
			res:  "system",
			want: true,
		},
		{
			sub:  "user.admin",
			act:  "ACT::POST/WRITE",
			res:  "post.1",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%-11v want to do %-25v on %-10v", tt.sub, tt.act, tt.res), func(t *testing.T) {
			result, _, err := e.EnforceEx(tt.sub, tt.act, tt.res)
			if err != nil {
				t.Error(err)
			}
			if result != tt.want {
				t.Errorf("EnforceEx() = %v, want %v", result, tt.want)
			}
		})
	}
}

func testRuleRaw(t *testing.T) {
	e := initCasbin()

	testRule := [][]string{
		{"user.1", "ACT::POST/WRITE", "post.1"},
		{"user.1", "ACT::POST/WRITE", "post.2"},
		{"user.2", "ACT::POST/WRITE", "post.3"},
		{"user.2", "ACT::POST/WRITE", "post.1"},
		{"user.2", "ACT::POST/READ", "post.1"},
		{"user.admin", "ACT::POST/WRITE", "post.1"},
		{"user.anon", "ACT::POST/READ", "post.1"},
		{"user.anon", "ACT::POST/WRITE", "post.1"},
		{"user.anon", "ACT::COMMENT/WRITE", "post.1"},
		{"user.anon", "ACT::COMMENT/READ", "comment.1"},
		{"user.admin", "ACT::USER_ADMIN/WRITE", "system"},
		{"user.admin", "ACT::POST/WRITE", "post.1"},
	}

	for _, v := range testRule {
		sub, act, res := v[0], v[1], v[2]
		result, reason, err := e.EnforceEx(sub, act, res)
		fmt.Printf("[%-5v] %-11v,%-25v,%-10v = %v ,err=%v\n", result, sub, act, res, reason, err)
	}
}

func testPattern(t *testing.T) {
	fmt.Println(util.KeyMatch("*", "post.1"))
	fmt.Println(util.KeyMatch("post.1", "*"))
	fmt.Println(util.GlobMatch("*", "post.1"))
	fmt.Println(util.GlobMatch("post.1", "*"))
}
