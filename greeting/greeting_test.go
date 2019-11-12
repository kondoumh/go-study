package greeting

import "testing"

func TestHello(t *testing.T) {
	want := "こんにちは世界。"
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
