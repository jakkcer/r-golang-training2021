package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

func TestStreamingDecode(t *testing.T) {
	type Book struct {
		Title, Author string
	}
	book := Book{"Point Counterpoint", "Aldous Huxley"}
	data, err := Marshal(book)
	if err != nil {
		t.Errorf("setting up test: %s", err)
		return
	}
	data = bytes.Repeat(data, 2)
	t.Logf("%s", data)

	dec := NewDecoder(bytes.NewReader(data))
	books := make([]Book, 0)
	for dec.More() {
		var b Book
		err := dec.Decode(&b)
		if err != nil {
			t.Errorf("Error after decoding %s: %s", books, err)
			return
		}
		books = append(books, b)
	}
	want := []Book{book, book}
	if !reflect.DeepEqual(want, books) {
		t.Errorf("Got %s, want %s", books, want)
	}
}
