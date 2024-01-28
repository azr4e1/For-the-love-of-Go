package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "122345678"
	cc, err := creditcard.New(want)

	if err != nil {
		t.Fatal("expected nil, got error")
	}
	got := cc.Number()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewFail(t *testing.T) {
	t.Parallel()
	number := ""
	_, err := creditcard.New(number)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
