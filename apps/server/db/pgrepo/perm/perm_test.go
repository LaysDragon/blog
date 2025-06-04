package perm

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

func main() {
	// Initialize a Xorm adapter with MySQL database.
	// a, err := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
	// if err != nil {
	// 	log.Fatalf("error: adapter: %s", err)
	// }

	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch)
	e.AddNamedMatchingFunc("g2", "KeyMatch2", util.KeyMatch)

	fmt.Println(e.EnforceEx("user.1", "ACT::POST/WRITE", "post.1"))
	fmt.Println(e.EnforceEx("user.1", "ACT::POST/WRITE", "post.2"))
	fmt.Println(e.EnforceEx("user.2", "ACT::POST/WRITE", "post.3"))
	fmt.Println(e.EnforceEx("user.2", "ACT::POST/WRITE", "post.1"))
	fmt.Println(e.EnforceEx("user.2", "ACT::POST/READ", "post.1"))
	fmt.Println(e.EnforceEx("admin", "ACT::POST/WRITE", "post.1"))
	fmt.Println(e.EnforceEx("user.anon", "ACT::POST/READ", "post.1"))
	fmt.Println(e.EnforceEx("user.anon", "ACT::POST/WRITE", "post.1"))
	fmt.Println(e.EnforceEx("user.anon", "ACT::COMMENT/WRITE", "post.1"))
	fmt.Println(e.EnforceEx("user.anon", "ACT::COMMENT/REAED", "comment.1"))

}
