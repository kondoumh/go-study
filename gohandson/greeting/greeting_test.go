package greeting

import (
	"bytes"
	"testing"
	"time"
	"golang.org/x/text/language"
)

func TestGreeting_Do(t *testing.T) {
	orgLang := lang
	lang = language.Japanese
	defer func() {
		lang = orgLang
	}()

	g := Greeting{
		Clock: ClockFunc(func() time.Time {
			return time.Date(2018, 8, 31, 06, 0, 0, 0, time.Local)
		}),
	}
	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error", err)
	}
	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}