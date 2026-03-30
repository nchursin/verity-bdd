package verity_test

import (
	"context"
	"testing"

	verity "github.com/nchursin/verity-bdd"
)

func TestRootAPIContractCompiles(t *testing.T) {
	var _ verity.Activity = verity.Do("noop", func(actor verity.Actor, ctx context.Context) error {
		return nil
	})

	var _ verity.Task = verity.TaskWhere("empty")

	q := verity.QuestionAbout("answer", func(actor verity.Actor, ctx context.Context) (int, error) {
		return 42, nil
	})

	_, _ = q.AnsweredBy(nil, context.Background())
}
