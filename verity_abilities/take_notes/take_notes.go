package take_notes

import (
	verity "github.com/nchursin/verity-bdd"
	internalnotes "github.com/nchursin/verity-bdd/internal/abilities/take_notes"
)

type TakeNotesAbility = internalnotes.TakeNotesAbility
type NoteBook = internalnotes.NoteBook

var UsingEmptyNotepad = internalnotes.UsingEmptyNotepad
var Using = internalnotes.Using
var NewNoteBook = internalnotes.NewNoteBook
var NotepadWith = internalnotes.NotepadWith

var TakeNoteOf = internalnotes.TakeNoteOf
var NoteValue = internalnotes.NoteValue

func Note[T any](key string) verity.Question[T] {
	return internalnotes.Note[T](key)
}
