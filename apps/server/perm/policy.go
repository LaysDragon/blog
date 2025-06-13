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
	if user.Role == domain.AdminRole {
		//p,user.admin,ROLE::ADMIN,*
		_, err := p.perm.AddPerm(User(user.ID), ROLE_ADMIN, ResWild())
		if err != nil {
			return err
		}
	} else {
		//p,user.1,ROLE::USER,*
		_, err := p.perm.AddPerm(User(user.ID), ROLE_USER, ResWild())
		if err != nil {
			return err
		}
	}
	//p,user.1,ROLE::USER/OWNER,user.2
	_, err := p.perm.AddPerm(User(user.ID), ROLE_OWNER, User(user.ID))
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
