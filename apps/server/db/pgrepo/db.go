package pgrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/LaysDragon/blog/apps/server/domain"
	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/volatiletech/sqlboiler/v4/boil"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CommonDb[T usecase.CommonRepo[T]] struct {
	db      boil.ContextExecutor
	builder func(db boil.ContextExecutor) T
}

func nilType[T any]() T {
	var t T
	return t
}

func (c *CommonDb[T]) BeginTx(ctx context.Context) (T, error) {
	if db, ok := c.db.(*sql.DB); ok {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return nilType[T](), err
		}
		return c.builder(tx), nil
	}
	return nilType[T](), errors.New("not *sql.DB")
}

func (c *CommonDb[T]) Commit() error {
	if tx, ok := c.db.(*sql.Tx); ok {
		return tx.Commit()
	}
	return errors.New("not *sql.Tx")
}

func (c *CommonDb[T]) Rollback() error {
	if tx, ok := c.db.(*sql.Tx); ok {
		return tx.Rollback()
	}
	return errors.New("not *sql.Tx")
}

type PostDb struct {
	CommonDb[usecase.PostRepo]
}

func NewPost(db boil.ContextExecutor) usecase.PostRepo {
	a := NewPost
	return &PostDb{
		CommonDb: CommonDb[usecase.PostRepo]{
			db:      db,
			builder: a,
		},
	}
}

func (r *PostDb) ToDb(post *domain.Post) *models.Post {
	if post == nil {
		return nil
	}
	return &models.Post{
		ID:        post.Id,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		SiteID:    post.SiteId,
		Content:   post.Content,
	}
}

func (r *PostDb) ToDomain(post *models.Post) *domain.Post {
	if post == nil {
		return nil
	}
	return &domain.Post{
		Id:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		SiteId:    post.SiteID,
		Content:   post.Content,
	}
}

func (r *PostDb) ById(ctx context.Context, id int) (*domain.Post, error) {
	post, err := models.FindPost(ctx, r.db, id)

	if err != nil {
		return nil, err
	}
	return r.ToDomain(post), nil
}

func (r *PostDb) Upsert(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	dbPost := r.ToDb(post)
	err := dbPost.Upsert(ctx, r.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, err
	}
	return r.ToDomain(dbPost), nil
}

func (r *PostDb) Delete(ctx context.Context, id int) error {
	_, err := models.Posts(models.PostWhere.ID.EQ(id)).DeleteAll(ctx, r.db, false)
	return err
}
