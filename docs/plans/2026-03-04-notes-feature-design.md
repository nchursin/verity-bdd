# Notes Feature Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Добавить Notes в стиле SerenityJS: `TakeNoteOf(value).As(key)` и `Note[T](key)` с хранением у актора и видимостью в отчётах.

**Architecture:** Новый ability `notes.NoteBook` с потокобезопасным `map[string]any`; Activity `TakeNoteOf` записывает, Question `Note[T]` читает. Шаги используют существующий `ActivityTracker` для репортинга через описания `#actor takes note "<key>"`. Интерфейсы Reporter не меняем.

**Tech Stack:** Go 1.22, TDD с `go test`.

---

### Task 1: Создать ability NoteBook (storage)

**Files:**
- Add: `serenity/abilities/notes/notebook.go`
- Tests: `serenity/abilities/notes/notebook_test.go`

**Steps:**
1. Написать красный тест: создание `NoteBook`, сохранение и чтение значений разных типов, overwrite того же ключа, потокобезопасность через параллельные goroutines (нет гонок, сравнение итогов).
2. Запустить `go test ./serenity/abilities/notes -run TestNoteBook -v` (ожидаем фейл).
3. Реализовать `NoteBook` с `map[string]any` + `sync.RWMutex`, методы `Set`, `Get`, `All`, ошибки при отсутствии ключа.
4. Запустить тесты, ожидаем PASS.

### Task 2: Activity TakeNoteOf(...).As(key)

**Files:**
- Add: `serenity/abilities/notes/take_note.go`
- Tests: `serenity/abilities/notes/take_note_test.go`

**Steps:**
1. Красный тест: актор с ability `TakeNotes()`; выполнение `TakeNoteOf("value").As("foo")` записывает в NoteBook; описание шага `#actor takes note "foo"`; ошибка, если нет ability NoteBook.
2. `go test ./serenity/abilities/notes -run TestTakeNote -v`.
3. Реализация: конструктор `TakeNotes()` возвращает ability (обёртка NoteBook), Activity `TakeNote` реализует `core.Activity`, в `PerformAs` ищет ability, кладёт значение; `Description()` возвращает `#actor takes note "<key>"`, `FailureMode()` — `core.FailFast`.
4. Прогон тестов.

### Task 3: Question Note[T](key) (Recall)

**Files:**
- Add: `serenity/abilities/notes/note_question.go`
- Tests: `serenity/abilities/notes/note_question_test.go`

**Steps:**
1. Красный тест: `Note[string]("foo").AnsweredBy(actor, ctx)` возвращает записанное значение; ошибка при отсутствии key; ошибка при неверном типе; необобщённый хелпер `NoteValue(key)` возвращает `any` без type assertion.
2. `go test ./serenity/abilities/notes -run TestNoteQuestion -v`.
3. Реализация: вопрос реализует `core.Question[T]`, достаёт ability NoteBook, делает type assertion, формирует описания `the note "<key>"`.
4. Прогон тестов.

### Task 4: Репортинг шага заметки

**Files:**
- Modify: (только тест) `serenity/abilities/notes/take_note_test.go` или новый `reporting_integration_test.go`
- Use: `serenity/reporting/mocks/mock_reporter.go`

**Steps:**
1. Красный тест: создать mock reporter, актор с `TestRunnerAdapter`, выполнить `TakeNoteOf("v").As("k")`; убедиться, что `OnStepStart` получил строку с `takes note "k"` и `OnStepFinish` вызван.
2. `go test ./serenity/abilities/notes -run TestTakeNoteReporting -v`.
3. При необходимости подправить описание шага (Activity `Description()`) или конструктор трекера, но без изменения интерфейсов Reporter.
4. Прогон тестов.

### Task 5: Документация и пример использования

**Files:**
- Add: `docs/examples/notes.md` (минимальный пример)
- Modify: `serenity/doc.go` или README, если есть секция про abilities.

**Steps:**
1. Написать пример: актор создаётся, `WhoCan(notes.TakeNotes())`, `AttemptsTo(notes.TakeNoteOf("token").As("auth"))`, далее `ensure.That(Note[string]("auth"), …)`.
2. Проверить сборку дока: `go test ./...` (или `go vet` если нужно), убедиться, что ничего не ломается.

### Task 6: Финальный прогон

**Steps:**
1. `go test ./...` в корне.
2. Подготовить commit (если потребуется) с сообщением в стиле conventional commits, например `feat: add actor notes ability`.
