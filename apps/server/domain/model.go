package domain

import "time"

type Account struct {
	Id         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Username   string
	Role       string
	Email      string
	PasswdHash string
}

type Site struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type Post struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	SiteId    int
	Content   string
}

type Attachtment struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	SiteId    int
	RelatedId int
	Url       string
}

type Comment struct {
	Id        int
	CreatedAt time.Time
	PostId    time.Time
	Email     string
	Name      string
	Content   string
}

type AccessLog struct {
	Id        int
	CreatedAt time.Time
	UserId    int
	Method    string
}
