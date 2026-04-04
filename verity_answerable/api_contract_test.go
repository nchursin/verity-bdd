package verity_answerable_test

import (
	"context"
	"testing"

	verity "github.com/nchursin/verity-bdd"
	answerable "github.com/nchursin/verity-bdd/verity_answerable"
)

func TestAnswerableAPIContractCompiles(t *testing.T) {
	_ = answerable.ValueOf(42)
	_ = answerable.ResultOf("description", func(ctx context.Context, actor verity.Actor) (int, error) {
		return 42, nil
	})
}
