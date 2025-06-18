package usecase

import (
	"context"
	"fmt"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (s *Post) Create(ctx context.Context, op perm.ResId, post *domain.Post) (*domain.Post, error) {
	err := s.perm.CheckE(op, perm.ACT_WRITE, perm.Site(post.SiteID))
	if err != nil {
		return nil, fmt.Errorf("create post failed:%w", err)
	}

	return s.repo.Upsert(ctx, post)
}
