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
	ACT_WRITE  ActStr = "WRITE"
	ACT_READ   ActStr = "READ"
	ACT_LIST   ActStr = "LIST"
	ACT_DELETE ActStr = "DELETE"

	ACT_WRITE_USER_ADMIN ActStr = "ACT::USER_ADMIN/WRITE"
	ACT_WRITE_SITE_OWNER ActStr = "ACT::SITE_OWNER/WRITE"

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

// ACT::{res_type}/{act}

func (a ActStr) Res(res ResStr) ActStr {
	if a.IsOverride() {
		return a
	}
	return ActStr(fmt.Sprintf("ACT::%v/%v", strings.ToUpper(string(res)), a))
}

func (a ActStr) Str() string {
	if a.IsOverride() {
		return string(a)
	}
	return fmt.Sprintf("ACT::%v", strings.ToUpper(string(a)))
}

// TODO: test ResStr
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
	// for user.system, record the source of the system op,for some internal op that will need to temp bypass perm check
	Source *ResId
}

// {res_type}.{id}
func (r ResId) Str() string {
	if r.ID == "" {
		return string(r.Name)
	}
	return fmt.Sprintf("%v.%v", r.Name, r.ID)
}

func (r ResId) String() string {
	if r.Source == nil {
		return r.Str()
	}
	return fmt.Sprintf("%v(%v)", r.Str(), r.Source.String())

}

func (r ResId) Type() ResStr {
	return r.Name
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

// user.system, for special internal logic that bypass all operation
func UserSystem(op ResId) ResId {
	if op.Name == RES_USER && op.ID == "system" {
		return op
	}
	return ResId{
		Name:   RES_USER,
		ID:     "system",
		Source: &op,
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

func SiteWild() ResId {
	return ResId{
		Name: RES_SITE,
		ID:   "*",
	}
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

func ResWild() ResId {
	return ResId{
		Name: "*",
		ID:   "",
	}
}
