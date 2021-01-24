package syncutil

import "testing"

func TestIncrement(t *testing.T) {
	c := &Counter{
		Name: "Access",
	}
	want := 1
	if got := c.Increment(); got != want {
		t.Errorf("c.Increment() = %d, want %d", got, want)
	}
}

func TestView(t *testing.T) {
	c := &Counter{
		Name: "Access",
	}
	c.Increment()
	c.Increment()
	want := 2
	if got := c.View(); got != want {
		t.Errorf("c.Increment() = %d, want %d", got, want)
	}
}
