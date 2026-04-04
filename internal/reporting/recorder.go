package reporting

import (
	"context"
	"sync"
)

type recorderContextKey struct{}

// StepRecorder collects attachments produced during a single activity step.
type StepRecorder struct {
	mu          sync.Mutex
	attachments []Attachment
}

// Collect adds an attachment to the recorder.
func (r *StepRecorder) Collect(att Attachment) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.attachments = append(r.attachments, att)
}

// Drain returns all collected attachments and resets the list.
func (r *StepRecorder) Drain() []Attachment {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := r.attachments
	r.attachments = nil
	return out
}

// ContextWithRecorder returns a child context with a fresh StepRecorder injected.
func ContextWithRecorder(ctx context.Context) (context.Context, *StepRecorder) {
	r := &StepRecorder{}
	return context.WithValue(ctx, recorderContextKey{}, r), r
}

// CollectAttachment adds att to the StepRecorder in ctx, if present.
// Safe to call with a context that has no recorder (no-op).
func CollectAttachment(ctx context.Context, att Attachment) {
	if r, ok := ctx.Value(recorderContextKey{}).(*StepRecorder); ok {
		r.Collect(att)
	}
}
