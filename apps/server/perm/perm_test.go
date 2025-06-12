package perm

import (
	"fmt"
	"log"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

func TestRule(t *testing.T) {
	// Initialize a Xorm adapter with MySQL database.
	// a, err := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
	// if err != nil {
	// 	log.Fatalf("error: adapter: %s", err)
	// }

	e, err := casbin.NewEnforcer("model.conf", "test_policy.csv")
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch)
	e.AddNamedMatchingFunc("g2", "KeyMatch2", util.KeyMatch)

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

	// fmt.Println(e.EnforceEx("user.1", "ACT::POST/WRITE", "post.1"))
	// fmt.Println(e.EnforceEx("user.1", "ACT::POST/WRITE", "post.2"))
	// fmt.Println(e.EnforceEx("user.2", "ACT::POST/WRITE", "post.3"))
	// fmt.Println(e.EnforceEx("user.2", "ACT::POST/WRITE", "post.1"))
	// fmt.Println(e.EnforceEx("user.2", "ACT::POST/READ", "post.1"))
	// fmt.Println(e.EnforceEx("user.admin", "ACT::POST/WRITE", "post.1"))
	// fmt.Println(e.EnforceEx("user.anon", "ACT::POST/READ", "post.1"))
	// fmt.Println(e.EnforceEx("user.anon", "ACT::POST/WRITE", "post.1"))
	// fmt.Println(e.EnforceEx("user.anon", "ACT::COMMENT/WRITE", "post.1"))
	// fmt.Println(e.EnforceEx("user.anon", "ACT::COMMENT/READ", "comment.1"))
	// fmt.Println(e.EnforceEx("user.admin", "ACT::USER_ADMIN/WRITE", "post.1"))
	// fmt.Println(e.EnforceEx("user.admin", "ACT::POST/WRITE", "post.1"))

}

func TestPattern(t *testing.T) {
	fmt.Println(util.KeyMatch("*", "post.1"))
	fmt.Println(util.KeyMatch("post.1", "*"))
	fmt.Println(util.GlobMatch("*", "post.1"))
	fmt.Println(util.GlobMatch("post.1", "*"))
}
