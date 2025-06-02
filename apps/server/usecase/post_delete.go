package usecase

import (
	"context"
)

func (s *Post) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
