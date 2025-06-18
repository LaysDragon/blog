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

	err = s.transactor.WithinTransaction(ctx, func(ctx context.Context) error {
		post, err = s.repo.Upsert(ctx, post)
		if err != nil {
			return fmt.Errorf("create post data failed:%w", err)
		}
		err = s.perm.Logic.AddPost(post)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}
