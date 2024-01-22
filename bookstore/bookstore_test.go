package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "For the Love of Go",
		Author: "John Arundel",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(book)
	got := result.Copies
	if err != nil {
		t.Fatalf("Not expected error value")
	}
	if got != want {
		t.Errorf("Want %d copies after buying 1 copy from a stock of %d copies, got %d", want, book.Copies, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("Want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	got := bookstore.GetAllBooks(catalog)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	want := bookstore.Book{
		Title: "For the Love of Go",
		ID:    1,
	}
	catalog := []bookstore.Book{
		{Title: "The Hunger Games",
			ID: 2},
		want,
	}
	got, _ := bookstore.GetBook(catalog, 1)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
