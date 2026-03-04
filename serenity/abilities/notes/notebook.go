package notes

import (
	"fmt"
	"sync"

	"github.com/nchursin/serenity-go/serenity/abilities"
)

// TakeNotesAbility wraps the NoteBook so it can be registered as an ability.
type TakeNotesAbility struct {
	*NoteBook
}

// TakeNotes returns a new ability instance that stores notes for an actor.
func TakeNotes() abilities.Ability {
	return &TakeNotesAbility{NoteBook: NewNoteBook()}
}

// NoteBook stores actor notes in a threadsafe map.
// It is meant to be used as an ability attached to an actor.
type NoteBook struct {
	mutex sync.RWMutex
	notes map[string]any
}

// NewNoteBook creates a new empty NoteBook.
func NewNoteBook() *NoteBook {
	return &NoteBook{notes: make(map[string]any)}
}

// Set saves a value under the provided key.
func (n *NoteBook) Set(key string, value any) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.notes[key] = value
}

// Get retrieves a value stored under the provided key.
// Returns an error when the key does not exist.
func (n *NoteBook) Get(key string) (any, error) {
	n.mutex.RLock()
	defer n.mutex.RUnlock()

	value, ok := n.notes[key]
	if !ok {
		return nil, fmt.Errorf("note %q not found", key)
	}

	return value, nil
}

// All returns a copy of all notes.
func (n *NoteBook) All() map[string]any {
	n.mutex.RLock()
	defer n.mutex.RUnlock()

	copy := make(map[string]any, len(n.notes))
	for key, value := range n.notes {
		copy[key] = value
	}

	return copy
}
