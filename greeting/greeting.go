package greeting

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Greet() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
