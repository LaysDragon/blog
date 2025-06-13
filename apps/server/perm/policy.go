package perm

import (
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
		return err
	}
	return nil
}

// func (p *PolicyLogic) AddSite(post *domain.Post) {
// 	p.perm.AddResRelation(Site(post.SiteID), Post(post.ID))
// }

func (p *PolicyLogic) AddPost(post *domain.Post) {
	p.perm.AddResRelation(Site(post.SiteID), Post(post.ID))
}

func (p *PolicyLogic) AddComment(comment *domain.Comment) {
	p.perm.AddResRelation(Post(comment.PostID), Comment(comment.ID))
}
