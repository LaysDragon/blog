package domain

import (
	"slices"
	"time"
)

type Enum interface {
	IsValid() bool
}

type AccountRole string

const (
	AdminRole AccountRole = "ROLE::ADMIN"
	UserRole  AccountRole = "ROLE::USER"
)

func (r AccountRole) IsValid() bool {
	return slices.Contains([]AccountRole{AdminRole, UserRole}, r)
}

type Account struct {
	ID         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Username   string
	Role       AccountRole
	Email      string
	PasswdHash string
}

type Site struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type Post struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	SiteID    int
	Content   string
}

type Attachtment struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	SiteID    int
	RelatedID int
	Url       string
}

type Comment struct {
	ID        int
	CreatedAt time.Time
	PostID    time.Time
	Email     string
	Name      string
	Content   string
}

type AccessLog struct {
	ID        int
	CreatedAt time.Time
	UserID    int
	Method    string
}
