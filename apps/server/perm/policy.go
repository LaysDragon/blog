package perm

import (
	"fmt"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
)

type PolicyLogic struct {
	enforcer *casbin.Enforcer
	log      *zap.Logger
	perm     *Perm
}

func (p *PolicyLogic) AddAccount(user *domain.Account) error {
	ps := &Polices{}

	if user.Role == domain.AdminRole {
		//p,user.admin,ROLE::ADMIN,*
		ps.AddPerm(User(user.ID), ROLE_ADMIN, ResWild())
	} else {
		//p,user.1,ROLE::USER,*
		ps.AddPerm(User(user.ID), ROLE_USER, ResWild())
	}
	//p,user.1,ROLE::USER/OWNER,user.2
	ps.AddPerm(User(user.ID), ROLE_OWNER, User(user.ID))

	_, err := p.perm.AddPolicies(ps)
	if err != nil {
		return fmt.Errorf("add account perm rule failed:%w", err)
	}
	return nil
}

func (p *PolicyLogic) AddSite(site *domain.Site, user *domain.Account) error {
	ps := &Polices{}
	p.perm.AddResRelation(User(user.ID), Site(site.ID))
	p.perm.AddPerm(User(user.ID), ROLE_OWNER, Site(site.ID))
	_, err := p.perm.AddPolicies(ps)
	if err != nil {
		return fmt.Errorf("add site perm rule failed:%w", err)
	}
	return nil
}

func (p *PolicyLogic) AddPost(post *domain.Post) error {
	_, err := p.perm.AddResRelation(Site(post.SiteID), Post(post.ID))
	if err != nil {
		return fmt.Errorf("add site perm rule failed:%w", err)
	}
	return nil
}

func (p *PolicyLogic) AddComment(comment *domain.Comment) {
	p.perm.AddResRelation(Post(comment.PostID), Comment(comment.ID))
}
