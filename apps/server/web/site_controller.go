package web

import (
	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"go.uber.org/zap"
)

// Site DTO
type Site struct {
	ID   int
	Name string
}

type SiteController struct {
	usecase *usecase.Site
	log     *zap.Logger
	perm    *perm.Perm
}

func NewSiteController(usecase *usecase.Site, log *zap.Logger, perm *perm.Perm) *SiteController {
	return &SiteController{usecase, log, perm}
}

func (c *SiteController) ToDto(site *domain.Site) *Site {
	return &Site{
		ID:   site.ID,
		Name: site.Name,
	}

}
