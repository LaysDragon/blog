package domain

import "time"

type Account struct {
	Id          int
	CreatedDate time.Time
	UpdatedDate time.Time
	Username    string
	Role        string
	Email       string
	PasswdHash  string
}

type Site struct {
	Id          int
	CreatedDate time.Time
	UpdatedDate time.Time
	Name        string
}

type Post struct {
	Id          int
	CreatedDate time.Time
	UpdatedDate time.Time
	SiteId      int
	Content     string
}

type Attachtment struct {
	Id          int
	CreatedDate time.Time
	UpdatedDate time.Time
	SiteId      int
	RelatedId   int
	Url         string
}

type Comment struct {
	Id          int
	CreatedDate time.Time
	PostId      time.Time
	Email       string
	Name        string
	Content     string
}

type AccessLog struct {
	Id        int
	Timestamp time.Time
	UserId    int
	Method    string
}
