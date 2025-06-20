package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/Thiht/transactor"
)

type PostRepo interface {
	CommonRepo[PostRepo]
	ById(ctx context.Context, id int) (*domain.Post, error)
	Upsert(ctx context.Context, post *domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, offset int, limit int, sid int) ([]*domain.Post, error)
}

type Post struct {
	repo       PostRepo
	transactor transactor.Transactor
	perm       *perm.Perm
}

func NewPost(repo PostRepo, transactor transactor.Transactor, perm *perm.Perm) *Post {
	return &Post{repo, transactor, perm}
}
