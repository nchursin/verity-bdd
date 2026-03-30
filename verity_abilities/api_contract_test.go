package verity_abilities_test

import (
	"testing"

	verity "github.com/nchursin/verity-bdd"
	"github.com/nchursin/verity-bdd/verity_abilities/api"
	"github.com/nchursin/verity-bdd/verity_abilities/take_notes"
)

func TestAbilitiesAPIContractCompiles(t *testing.T) {
	var _ verity.Ability = api.CallAnApiAt("https://example.com")

	notebook := take_notes.NewNoteBook()
	_ = take_notes.Using(notebook)
	_ = take_notes.UsingEmptyNotepad()
}
