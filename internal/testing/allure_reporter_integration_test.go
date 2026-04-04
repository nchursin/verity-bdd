package testing

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nchursin/verity-bdd/internal/abilities/take_notes"
	"github.com/nchursin/verity-bdd/internal/core"
	"github.com/nchursin/verity-bdd/internal/reporting/allure_reporter"
)

func TestVerityTest_WithAllureReporter_WritesResults(t *testing.T) {
	t.Parallel()

	resultsDir := t.TempDir()
	reporter := allure_reporter.NewAllureReporterWithDir(resultsDir)

	test := NewVerityTest(t, Scene{
		Context:  context.Background(),
		Reporter: reporter,
	})

	actor := test.ActorCalled("Sam").WhoCan(take_notes.UsingEmptyNotepad())
	ability, err := actor.AbilityTo(&take_notes.TakeNotesAbility{})
	require.NoError(t, err)
	notebook := ability.(*take_notes.TakeNotesAbility)
	notebook.Set("token", "secret")

	actor.AttemptsTo(core.Do("records an action", func(ctx context.Context, actor core.Actor) error {
		return nil
	}))

	test.Shutdown()

	resultPath := readResultFilePath(t, resultsDir)
	payload, err := os.ReadFile(resultPath)
	require.NoError(t, err)

	var result map[string]any
	require.NoError(t, json.Unmarshal(payload, &result))

	steps, ok := result["steps"].([]any)
	require.True(t, ok)
	require.Len(t, steps, 1)

	attachments, ok := result["attachments"].([]any)
	require.True(t, ok)
	require.Len(t, attachments, 1)

	attachment := attachments[0].(map[string]any)
	require.Equal(t, "notes", attachment["name"])

	source := attachment["source"].(string)
	attachmentPayload, err := os.ReadFile(filepath.Join(resultsDir, source))
	require.NoError(t, err)

	var notesJSON map[string]map[string]any
	require.NoError(t, json.Unmarshal(attachmentPayload, &notesJSON))
	require.Equal(t, "secret", notesJSON["Sam"]["token"])
}

func readResultFilePath(t *testing.T, dir string) string {
	t.Helper()

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)

	results := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		results = append(results, filepath.Join(dir, entry.Name()))
	}

	require.Len(t, results, 1)
	return results[0]
}
