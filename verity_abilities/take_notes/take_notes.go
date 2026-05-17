package take_notes

import (
	verity "github.com/verity-bdd/verity-bdd"
	internalnotes "github.com/verity-bdd/verity-bdd/internal/abilities/take_notes"
)

// TakeNotesAbility wraps the NoteBook so it can be registered as an ability.
type TakeNotesAbility = internalnotes.TakeNotesAbility

// NoteBook stores actor notes in a thread-safe map.
// It is meant to be used as an ability attached to an actor.
type NoteBook = internalnotes.NoteBook

// UsingEmptyNotepad returns a new ability instance with an empty notepad.
var UsingEmptyNotepad = internalnotes.UsingEmptyNotepad

// Using returns an ability that stores notes in the provided notepad.
var Using = internalnotes.Using

// NewNoteBook creates a new empty NoteBook.
var NewNoteBook = internalnotes.NewNoteBook

// NotepadWith creates a notepad pre-filled with the provided values.
var NotepadWith = internalnotes.NotepadWith

// TakeNoteOf starts a TakeNote activity definition for the given value.
var TakeNoteOf = internalnotes.TakeNoteOf

// NoteValue returns an untyped question that retrieves the note stored under the given key.
var NoteValue = internalnotes.NoteValue

// Note returns a typed question that retrieves the note stored under the given key as type T.
func Note[T any](key string) verity.Question[T] {
	return internalnotes.Note[T](key)
}
