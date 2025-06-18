package pgrepo

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/LaysDragon/blog/apps/server/domain"
	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/utils"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PostDb struct {
	CommonDb[usecase.PostRepo]
}

func NewPost(db stdlibTransactor.DBGetter) usecase.PostRepo {
	return &PostDb{
		CommonDb: CommonDb[usecase.PostRepo]{
			db:      db,
			builder: NewPost,
		},
	}
}

func (r *PostDb) ToDb(post *domain.Post) *models.Post {
	if post == nil {
		return nil
	}
	return &models.Post{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		SiteID:    post.SiteID,
		Content:   post.Content,
	}
}

func (r *PostDb) ToDomain(post *models.Post) *domain.Post {
	if post == nil {
		return nil
	}
	return &domain.Post{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		SiteID:    post.SiteID,
		Content:   post.Content,
	}
}

func (r *PostDb) ById(ctx context.Context, id int) (*domain.Post, error) {
	post, err := models.FindPost(ctx, r.db(ctx), id)

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(post), nil
}

func (r *PostDb) Upsert(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	dbPost := r.ToDb(post)
	err := dbPost.Upsert(ctx, r.db(ctx), true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(dbPost), nil
}

func (r *PostDb) Delete(ctx context.Context, id int) error {
	affRow, err := models.Posts(models.PostWhere.ID.EQ(id)).DeleteAll(ctx, r.db(ctx), false)
	if affRow == 0 {
		return usecase.ItemNotExistedError{}
	}
	return ErrorTranslate(err)
}

func (r *PostDb) List(ctx context.Context, offset int, limit int, sid int) ([]*domain.Post, error) {
	queries := []QueryMod{Offset(offset), Limit(limit), OrderBy(models.PostColumns.ID)}

	if sid > 0 {
		queries = append(queries, models.PostWhere.SiteID.EQ(sid))
	}

	posts, err := models.Posts(queries...).All(ctx, r.db(ctx))
	if err != nil {
		return nil, ErrorTranslate(err)
	}

	return utils.MappingFunc(posts, r.ToDomain), nil

}
