package perm

import (
	"fmt"
	"strconv"
	"strings"
)

type ActStr string
type ResStr string
type RoleStr string

const (
	// ACT_USER_WRITE  ActStr = "ACT::USER/WRITE"
	// ACT_USER_READ   ActStr = "ACT::USER/READ"
	// ACT_USER_DELETE ActStr = "ACT::USER/DELETE"

	// ACT_POST_WRITE  ActStr = "ACT::POST/WRITE"
	// ACT_POST_READ   ActStr = "ACT::POST/READ"
	// ACT_POST_DELETE ActStr = "ACT::POST/DELETE"

	// ACT_SITE_WRITE ActStr = "ACT::SITE/WRITE"
	// ACT_SITE_READ  ActStr = "ACT::SITE/READ"

	// ACT_COMMENT__WRITE  ActStr = "ACT::COMMENT/WRITE"
	// ACT_COMMENT__READ   ActStr = "ACT::COMMENT/READ"
	// ACT_COMMENT__DELETE ActStr = "ACT::COMMENT/DELETE"

	// ROLE_USER_OWNER RoleStr = "ROLE::USER/OWNER"
	// ROLE_SITE_OWNER RoleStr = "ROLE::SITE/OWNER"
	// ROLE_POST_OWNER RoleStr = "ROLE::POST/OWNER"
	ACT_WRITE  ActStr = "WRITE"
	ACT_READ   ActStr = "READ"
	ACT_DELETE ActStr = "DELETE"

	ACT_WRITE_USER_ADMIN ActStr = "ACT::USER_ADMIN/WRITE"

	//system implied top level scope
	RES_SYSTEM ResStr = "system"

	RES_USER    ResStr = "user"
	RES_POST    ResStr = "post"
	RES_SITE    ResStr = "site"
	RES_COMMENT ResStr = "comment"

	ROLE_OWNER RoleStr = "OWNER"

	ROLE_ADMIN RoleStr = "ROLE::ADMIN"
	ROLE_USER  RoleStr = "ROLE::USER"
)

// func (r RoleStr) Str() string {
// 	return fmt.Sprintf("ROLE::%v", r)
// }

func (a RoleStr) IsOverride() bool {
	return strings.HasPrefix(string(a), "ROLE::")
}

func (a ActStr) IsOverride() bool {
	return strings.HasPrefix(string(a), "ACT::")
}

// {res_type}.{id}
func (r ResStr) ID(id int) ResId {
	return ResId{
		ID:   strconv.Itoa(id),
		Name: r,
	}
}

// // {res_type}.new
// func (r ResStr) New() ResId {
// 	return ResId{
// 		ID:   "new",
// 		Name: r,
// 	}
// }

// ACT::{res_type}/{act}
func (r ResStr) Act(act ActStr) string {
	if act.IsOverride() {
		return string(act)
	}
	return fmt.Sprintf("ACT::%v/%v", r, act)
}

// ROLE::{res_type}/{role}
func (r ResStr) Role(role RoleStr) string {
	if role.IsOverride() {
		return string(role)
	}
	return fmt.Sprintf("ROLE::%v/%v", strings.ToUpper(string(r)), role)
}

type ResId struct {
	ID   string
	Name ResStr
}

// {res_type}.{id}
func (r ResId) Str() string {
	if r.ID == "" {
		return string(r.Name)
	}
	return fmt.Sprintf("%v.%v", r.Name, r.ID)
}

func (r ResId) Type() ResStr {
	return r.Name
}

func ResWild() ResId {
	return ResId{
		Name: "*",
		ID:   "*",
	}
}

func User(id int) ResId {
	return RES_USER.ID(id)
}

// user.anon
func UserAnon() ResId {
	return ResId{
		Name: RES_USER,
		ID:   "anon",
	}
}

// // user.new_admin
// func UserNewAdmin() ResId {
// 	return ResId{
// 		Name: RES_USER,
// 		ID:   "new_admin",
// 	}
// }

func Site(id int) ResId {
	return RES_SITE.ID(id)
}
func Post(id int) ResId {
	return RES_POST.ID(id)
}
func Comment(id int) ResId {
	return RES_COMMENT.ID(id)
}

func System() ResId {
	return ResId{
		Name: RES_SYSTEM,
		ID:   "",
	}
}
