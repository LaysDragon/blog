package pgrepo

import (
	"context"
	"database/sql"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/service"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PostRepo struct {
	db *sql.DB
}

// type Post struct{}

func NewPostRepo(db *sql.DB) service.PostRepository {
	return &PostRepo{db}
}

func (repo *PostRepo) ById(id int) (*domain.Post, error) {
	post, err := models.Posts(Where("id = ?", id)).One(context.Background(), repo.db)
	if err != nil {
		return nil, err
	}
	return &domain.Post{
		Id:      post.ID,
		Content: post.Content,
	}, nil
}
