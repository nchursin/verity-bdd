package examples

import (
	"bytes"
	"errors"
	"fmt"

	reporting "github.com/nchursin/verity-bdd/verity_reporting"
	"github.com/nchursin/verity-bdd/verity_reporting/console_reporter"
)

type exampleResult struct {
	name        string
	status      reporting.Status
	duration    float64
	err         error
	attachments []reporting.Attachment
}

func (sr *exampleResult) Name() string {
	return sr.name
}

func (sr *exampleResult) Status() reporting.Status {
	return sr.status
}

func (sr *exampleResult) Duration() float64 {
	return sr.duration
}

func (sr *exampleResult) Error() error {
	return sr.err
}

func (sr *exampleResult) Attachments() []reporting.Attachment {
	return sr.attachments
}

// ExampleConsoleReporter_withAttachments shows how attachments are rendered in console output.
func ExampleConsoleReporter_withAttachments() {
	var buf bytes.Buffer
	reporter := console_reporter.NewConsoleReporter()
	reporter.SetOutput(&buf)

	reporter.OnTestFinish(&exampleResult{
		name:     "ExampleFailed",
		status:   reporting.StatusFailed,
		duration: 0.05,
		err:      errors.New("boom"),
		attachments: []reporting.Attachment{
			{
				Name:        "notes",
				ContentType: "application/json",
				Content:     []byte(`{"token":"secret"}`),
			},
		},
	})

	fmt.Print(buf.String())
	// Output:
	// ❌ ExampleFailed: FAILED (0.05s)
	//    Error: boom
	//    Attachments:
	//    - notes (application/json): {"token":"secret"}
	//
}
