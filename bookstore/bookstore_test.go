package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"sort"
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
	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go", ID: 1},
		2: {Title: "The Power of Go: Tools", ID: 2},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go", ID: 1},
		{Title: "The Power of Go: Tools", ID: 2},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool { return got[i].ID < got[j].ID })
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 1, Title: "The Power of Go: Tools"},
	}
	want := bookstore.Book{
		Title: "For the Love of Go",
		ID:    1,
	}
	got, err := catalog.GetBook(1)
	if err != nil {
		t.Fatal("Error not expected.")
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 1, Title: "The Power of Go: Tools"},
	}
	_, err := catalog.GetBook(3)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		PriceCents:      300,
		DiscountPercent: 30,
	}
	var want int = 210
	// got := bookstore.NetPriceCents(b)
	got := b.NetPriceCents()
	if got != want {
		t.Errorf("original price: %d, discount: %d, want %d but got %d", b.PriceCents, b.DiscountPercent, want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	var b bookstore.Book = bookstore.Book{
		PriceCents:      300,
		DiscountPercent: 30,
	}
	var newPrice int = 200
	// var ptr *bookstore.Book = &b
	err := b.SetPriceCents(newPrice)
	if err != nil {
		t.Fatal(err)
	}
	want := newPrice
	got := b.PriceCents
	if got != want {
		t.Errorf("want new price %d, got %d", want, got)
	}
}

func TestSetPriceCentsFail(t *testing.T) {
	t.Parallel()
	var price int = 300
	var b bookstore.Book = bookstore.Book{
		PriceCents:      price,
		DiscountPercent: 30,
	}
	var newPrice int = -1
	var err error = b.SetPriceCents(newPrice)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	want := price
	got := b.PriceCents
	if got != want {
		t.Errorf("want old price %d, got %d", want, got)
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	// test cases for categories
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}
	for _, want := range cats {
		err := b.SetCategory(want)
		if err != nil {
			t.Fatal(err)
		}
		got := b.Category()
		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	}
}

func TestSetCategoryFail(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory(-1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
