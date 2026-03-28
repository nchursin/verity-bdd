# Notes (заметки актора)

Минимальный пример использования заметок как в SerenityJS:

```go
package examples

import (
    "testing"

    "github.com/nchursin/serenity-go/serenity/abilities/take_notes"
    serenitytesting "github.com/nchursin/serenity-go/serenity/testing"
)

func TestNotesExample(t *testing.T) {
    test := serenitytesting.NewSerenityTest(t, serenitytesting.Scene{})
    actor := test.ActorCalled("Nina").WhoCan(take_notes.UsingEmptyNotepad())

    actor.AttemptsTo(
        take_notes.TakeNoteOf("Bearer abc123").As("auth token"),
    )

    token, ok := actor.AnswersTo(take_notes.Note[string]("auth token"))
    if !ok {
        t.Fatalf("note not found")
    }
    if token != "Bearer abc123" {
        t.Fatalf("unexpected token: %s", token)
    }
}
```

Что происходит:
- `UsingEmptyNotepad()` добавляет ability хранения заметок актору.
- `Using(NotepadWith(...))` позволяет задать стартовые заметки, например имя актора и роль.
- `TakeNoteOf(...).As("auth token")` записывает значение и создаёт шаг в отчёте вида `Nina takes note "auth token"`.
- `Note[string]("auth token")` безопасно читает заметку с проверкой типа.
