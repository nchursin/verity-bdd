package notes

import (
	"strconv"
	"sync"
	"testing"
)

func TestNoteBookStoresAndRetrievesValues(t *testing.T) {
	noteBook := NewNoteBook()

	noteBook.Set("greeting", "hello")
	noteBook.Set("count", 42)

	value, err := noteBook.Get("greeting")
	if err != nil {
		t.Fatalf("expected greeting note, got error: %v", err)
	}
	if value != "hello" {
		t.Fatalf("expected 'hello', got %v", value)
	}

	value, err = noteBook.Get("count")
	if err != nil {
		t.Fatalf("expected count note, got error: %v", err)
	}
	if value != 42 {
		t.Fatalf("expected 42, got %v", value)
	}
}

func TestNoteBookOverwritesValues(t *testing.T) {
	noteBook := NewNoteBook()

	noteBook.Set("key", "first")
	noteBook.Set("key", "second")

	value, err := noteBook.Get("key")
	if err != nil {
		t.Fatalf("expected note, got error: %v", err)
	}
	if value != "second" {
		t.Fatalf("expected overwritten value 'second', got %v", value)
	}
}

func TestNoteBookReturnsErrorWhenMissing(t *testing.T) {
	noteBook := NewNoteBook()

	_, err := noteBook.Get("missing")
	if err == nil {
		t.Fatalf("expected error for missing note")
	}
}

func TestNoteBookAllReturnsCopy(t *testing.T) {
	noteBook := NewNoteBook()
	noteBook.Set("k1", "v1")
	noteBook.Set("k2", "v2")

	all := noteBook.All()
	all["k1"] = "changed"

	value, _ := noteBook.Get("k1")
	if value != "v1" {
		t.Fatalf("expected original map to stay unchanged, got %v", value)
	}
}

func TestNoteBookIsConcurrentSafe(t *testing.T) {
	noteBook := NewNoteBook()
	count := 200
	wg := sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		i := i
		go func() {
			defer wg.Done()
			noteBook.Set("key"+strconv.Itoa(i), i)
		}()
	}

	wg.Wait()

	for i := 0; i < count; i++ {
		value, err := noteBook.Get("key" + strconv.Itoa(i))
		if err != nil {
			t.Fatalf("expected key %d, got error: %v", i, err)
		}
		if value != i {
			t.Fatalf("expected %d, got %v", i, value)
		}
	}
}
