// Code generated by SQLBoiler boilingseed-0.1.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package seeds

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	models "github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

var (
	postColumnsWithDefault = []string{"id", "created_at", "updated_at", "deleted_at"}
	postDBTypes            = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `SiteID`: `integer`, `Content`: `text`}
)

func defaultPostForeignKeySetter(i int, o *models.Post, allSites models.SiteSlice) error {
	if len(allSites) > 0 {
		// set site
		SiteKey := int(math.Mod(float64(i), float64(len(allSites))))
		site := allSites[SiteKey]

		o.SiteID = site.ID

	}
	return nil
}

// defaultRandomPost creates a random model.Post
// Used when RandomPost is not set in the Seeder
func defaultRandomPost() (*models.Post, error) {
	o := &models.Post{}
	seed := randomize.NewSeed()
	err := randomize.Struct(seed, o, postDBTypes, true, postColumnsWithDefault...)

	return o, err
}

func (s Seeder) seedPosts(ctx context.Context, exec boil.ContextExecutor) error {
	fmt.Println("Adding Posts")
	PostsToAdd := s.MinPostsToSeed

	randomFunc := s.RandomPost
	if randomFunc == nil {
		randomFunc = defaultRandomPost
	}

	fkFunc := s.PostForeignKeySetter
	if fkFunc == nil {
		fkFunc = defaultPostForeignKeySetter
	}

	sites, err := models.Sites().All(ctx, exec)
	if err != nil {
		return fmt.Errorf("error getting sites: %w", err)
	}

	if s.PostsPerSite*len(sites) > PostsToAdd {
		PostsToAdd = s.PostsPerSite * len(sites)
	}

	for i := 0; i < PostsToAdd; i++ {
		// create model
		o, err := randomFunc()
		if err != nil {
			return fmt.Errorf("unable to get Random Post: %w", err)
		}

		// Set foreign keys
		err = fkFunc(i, o, sites)
		if err != nil {
			return fmt.Errorf("unable to get set foreign keys for Post: %w", err)
		}

		// insert model
		if err := o.Insert(ctx, exec, boil.Infer()); err != nil {
			return fmt.Errorf("unable to insert Post: %w", err)
		}
	}

	// run afterAdd
	if s.AfterPostsAdded != nil {
		if err := s.AfterPostsAdded(ctx); err != nil {
			return fmt.Errorf("error running AfterPostsAdded: %w", err)
		}
	}

	fmt.Println("Finished adding Posts")
	return nil
}

// These packages are needed in SOME models
// This is to prevent errors in those that do not need it
var _ = math.E
var _ = queries.Query{}

// This is to force strconv to be used. Without it, it causes an error because strconv is imported by ALL the drivers
var _ = strconv.IntSize

// post is here to prevent erros due to driver "BasedOnType" imports.
type post struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
	SiteID    int
	Content   string
}
