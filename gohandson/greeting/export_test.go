package greeting

import "golang.org/x/text/language"

func ExportSetLang(l language.Tag) func() {
	orgLang := lang
	lang = l
	return func() {
		lang = orgLang
	}
}