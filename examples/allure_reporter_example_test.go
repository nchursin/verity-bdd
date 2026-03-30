package examples

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nchursin/verity-bdd/verity/abilities/take_notes"
	"github.com/nchursin/verity-bdd/verity/core"
	"github.com/nchursin/verity-bdd/verity/reporting/allure_reporter"
	verity "github.com/nchursin/verity-bdd/verity/testing"
)

// TestAllureReporterExample_GeneratesReportFiles is a documentation-style example:
// create an Allure reporter, run actor steps, then assert generated result files.
func TestAllureReporterExample_GeneratesReportFiles(t *testing.T) {
	resultsDir := t.TempDir()
	reporter := allure_reporter.NewAllureReporterWithDir(resultsDir)

	test := verity.NewVerityTest(t, verity.Scene{
		Context:  context.Background(),
		Reporter: reporter,
	})

	actor := test.ActorCalled("Sam").WhoCan(take_notes.UsingEmptyNotepad())
	actor.AttemptsTo(
		core.Do("#actor does something", func(actor core.Actor, ctx context.Context) error {
			return nil
		}),
		core.Do("#actor records an Allure-friendly step", func(actor core.Actor, ctx context.Context) error {
			ability, err := actor.AbilityTo(&take_notes.TakeNotesAbility{})
			if err != nil {
				return err
			}

			notebook := ability.(*take_notes.TakeNotesAbility)
			notebook.Set("token", "secret")
			return nil
		}),
		core.Do("#actor does something else", func(actor core.Actor, ctx context.Context) error {
			return nil
		}),
	)

	test.Shutdown()

	resultFile := findSingleAllureResultFile(t, resultsDir)
	resultPayload, err := os.ReadFile(resultFile)
	require.NoError(t, err)

	var result map[string]any
	require.NoError(t, json.Unmarshal(resultPayload, &result))

	require.Equal(t, "TestAllureReporterExample_GeneratesReportFiles", result["name"])
	require.Equal(t, "passed", result["status"])

	steps, ok := result["steps"].([]any)
	require.True(t, ok)
	require.Len(t, steps, 3)

	attachments, ok := result["attachments"].([]any)
	require.True(t, ok)
	require.Len(t, attachments, 1)

	notesAttachment, ok := attachments[0].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "notes", notesAttachment["name"])

	source, ok := notesAttachment["source"].(string)
	require.True(t, ok)
	require.NotEmpty(t, source)

	attachmentPayload, err := os.ReadFile(filepath.Join(resultsDir, source))
	require.NoError(t, err)

	var notes map[string]map[string]any
	require.NoError(t, json.Unmarshal(attachmentPayload, &notes))
	require.Equal(t, "secret", notes["Sam"]["token"])
}

func findSingleAllureResultFile(t *testing.T, resultsDir string) string {
	t.Helper()

	entries, err := os.ReadDir(resultsDir)
	require.NoError(t, err)

	resultFiles := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		resultFiles = append(resultFiles, filepath.Join(resultsDir, entry.Name()))
	}

	require.Len(t, resultFiles, 1)
	return resultFiles[0]
}
