package greeting

import (
	"errors"

	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Greet() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}

func DoMinus(arg int) int {
	res, _ := minus9(arg)
	return res
}

func minus9(num int) (int, error) {
	if num < 10 {
		return 0, errors.New("num shoud be 10 or more!")
	}
	return num - 9, nil
}
