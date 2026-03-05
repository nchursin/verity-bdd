package testing

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/nchursin/serenity-go/serenity/reporting"
	"github.com/nchursin/serenity-go/serenity/reporting/console_reporter"
	reportingMocks "github.com/nchursin/serenity-go/serenity/reporting/mocks"
	"github.com/nchursin/serenity-go/serenity/testing/mocks"

	"github.com/nchursin/serenity-go/serenity/abilities/notes"
)

func TestSerenityTestWithConsoleReporter(t *testing.T) {
	ctx := context.Background()
	// Create a SerenityTest with console reporter
	test := NewSerenityTestWithReporter(ctx, t, console_reporter.NewConsoleReporter())

	actor := test.ActorCalled("TestActor")
	require.NotNil(t, actor)

	// Verify that reporter is configured
	adapter := test.GetReporterAdapter()
	require.NotNil(t, adapter)
	require.IsType(t, &console_reporter.ConsoleReporter{}, adapter.GetReporter())
}

func TestNewSerenityTestUsesConsoleReporter(t *testing.T) {
	ctx := context.Background()
	test := NewSerenityTestWithContext(ctx, t)

	adapter := test.GetReporterAdapter()
	require.NotNil(t, adapter)

	// Verify it's a ConsoleReporter
	reporter := adapter.GetReporter()
	_, isConsole := reporter.(*console_reporter.ConsoleReporter)
	require.True(t, isConsole, "Expected ConsoleReporter")
}

func TestSerenityTestLifecycleReporting(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReporter := reportingMocks.NewMockReporter(ctrl)
	mockTestContext := mocks.NewMockTestContext(ctrl)

	// Expect test lifecycle events
	mockReporter.EXPECT().OnTestStart("TestExample")
	mockReporter.EXPECT().OnTestFinish(gomock.Any()).Do(func(result reporting.TestResult) {
		require.Equal(t, "TestExample", result.Name())
		require.Equal(t, reporting.StatusPassed, result.Status())
		require.True(t, result.Duration() >= 0)
		require.NoError(t, result.Error())
	})

	mockTestContext.EXPECT().Name().Return("TestExample")
	mockTestContext.EXPECT().Failed().Return(false)
	mockTestContext.EXPECT().Helper()
	mockTestContext.EXPECT().Cleanup(gomock.Any())

	ctx := context.Background()
	test := NewSerenityTestWithReporter(ctx, mockTestContext, mockReporter)

	// Simulate test end
	test.Shutdown()
}

func TestSerenityTestLifecycleReportingFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReporter := reportingMocks.NewMockReporter(ctrl)
	mockTestContext := mocks.NewMockTestContext(ctrl)

	// Expect test lifecycle events for failed test
	mockReporter.EXPECT().OnTestStart("FailedTest")
	mockReporter.EXPECT().OnTestFinish(gomock.Any()).Do(func(result reporting.TestResult) {
		require.Equal(t, "FailedTest", result.Name())
		require.Equal(t, reporting.StatusFailed, result.Status())
		require.True(t, result.Duration() >= 0)
		require.Error(t, result.Error())
		require.Equal(t, "test failed", result.Error().Error())
	})

	mockTestContext.EXPECT().Name().Return("FailedTest")
	mockTestContext.EXPECT().Failed().Return(true)
	mockTestContext.EXPECT().Helper()
	mockTestContext.EXPECT().Cleanup(gomock.Any())

	ctx := context.Background()
	test := NewSerenityTestWithReporter(ctx, mockTestContext, mockReporter)

	// Simulate test end
	test.Shutdown()
}

func TestSerenityTestAddsNotesAttachmentOnShutdown(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReporter := reportingMocks.NewMockReporter(ctrl)
	mockTestContext := mocks.NewMockTestContext(ctrl)

	mockReporter.EXPECT().OnTestStart("NotesTest")
	mockReporter.EXPECT().OnTestFinish(gomock.Any()).Do(func(result reporting.TestResult) {
		attachments := result.Attachments()
		require.Len(t, attachments, 1)

		attachment := attachments[0]
		require.Equal(t, "notes", attachment.Name)
		require.Equal(t, "application/json", attachment.ContentType)

		var payload map[string]map[string]any
		require.NoError(t, json.Unmarshal(attachment.Content, &payload))

		notesForActor, ok := payload["Sam"]
		require.True(t, ok, "expected notes for actor Sam")
		require.Equal(t, "secret", notesForActor["token"])
		intCount, ok := notesForActor["count"].(float64)
		require.True(t, ok, "expected numeric count")
		require.Equal(t, float64(2), intCount)
	})

	mockTestContext.EXPECT().Name().Return("NotesTest")
	mockTestContext.EXPECT().Failed().Return(false)
	mockTestContext.EXPECT().Helper()
	mockTestContext.EXPECT().Cleanup(gomock.Any())

	ctx := context.Background()
	test := NewSerenityTestWithReporter(ctx, mockTestContext, mockReporter)

	actor := test.ActorCalled("Sam").WhoCan(notes.TakeNotes())
	ability, err := actor.AbilityTo(&notes.TakeNotesAbility{})
	require.NoError(t, err)
	notebook := ability.(*notes.TakeNotesAbility)
	notebook.Set("token", "secret")
	notebook.Set("count", 2)

	test.Shutdown()
}

func TestSerenityTestAddsNotesAttachmentOnFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReporter := reportingMocks.NewMockReporter(ctrl)
	mockTestContext := mocks.NewMockTestContext(ctrl)

	mockReporter.EXPECT().OnTestStart("NotesFailedTest")
	mockReporter.EXPECT().OnTestFinish(gomock.Any()).Do(func(result reporting.TestResult) {
		require.Equal(t, "NotesFailedTest", result.Name())
		require.Equal(t, reporting.StatusFailed, result.Status())
		require.Error(t, result.Error())
		require.EqualError(t, result.Error(), "test failed")

		attachments := result.Attachments()
		require.Len(t, attachments, 1)

		attachment := attachments[0]
		require.Equal(t, "notes", attachment.Name)
		require.Equal(t, "application/json", attachment.ContentType)

		var payload map[string]map[string]any
		require.NoError(t, json.Unmarshal(attachment.Content, &payload))

		notesForActor, ok := payload["Sam"]
		require.True(t, ok, "expected notes for actor Sam")
		require.Equal(t, "secret", notesForActor["token"])
		intCount, ok := notesForActor["count"].(float64)
		require.True(t, ok, "expected numeric count")
		require.Equal(t, float64(3), intCount)
	})

	mockTestContext.EXPECT().Name().Return("NotesFailedTest")
	mockTestContext.EXPECT().Failed().Return(true)
	mockTestContext.EXPECT().Helper()
	mockTestContext.EXPECT().Cleanup(gomock.Any())

	ctx := context.Background()
	test := NewSerenityTestWithReporter(ctx, mockTestContext, mockReporter)

	actor := test.ActorCalled("Sam").WhoCan(notes.TakeNotes())
	ability, err := actor.AbilityTo(&notes.TakeNotesAbility{})
	require.NoError(t, err)
	notebook := ability.(*notes.TakeNotesAbility)
	notebook.Set("token", "secret")
	notebook.Set("count", 3)

	test.Shutdown()
}
