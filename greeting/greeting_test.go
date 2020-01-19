package greeting

import "testing"

func TestHello(t *testing.T) {
	want := "Ahoy world!"
	if got := Greet(); got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}

func TestDoMinus(t *testing.T) {
	want := 91
	if got := DoMinus(100); got != want {
		t.Errorf("DoMinus(100) = %d, want %d", got, want)
	}
}
