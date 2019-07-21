package calc

import (
	"testing"
	"fmt"
)

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Fatal("sum(1,2) shuoud be 3. but doesn't match")
	}
}

func ExampleHello() {
	fmt.Println(sum(1, 2))
	// Output: 3
}