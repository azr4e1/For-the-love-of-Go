package basics_test

import (
	"basics"
	"testing"
)

func TestWithdrawValid(t *testing.T) {
	t.Parallel()
	balance := 100
	amount := 20
	want := 80
	got, err := basics.Withdraw(balance, amount)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithdrawInvalid(t *testing.T) {
	t.Parallel()
	balance := 20
	amount := 100
	_, err := basics.Withdraw(balance, amount)
	if err == nil {
		t.Fatal("want error for invalid withdrawal, got nil")
	}
}

func TestApply(t *testing.T) {
	t.Parallel()
	want := 2
	got := basics.Apply(1, func(x int) int { return x * 2 })
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
