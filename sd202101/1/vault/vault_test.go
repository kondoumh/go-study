package vault

import (
	"fmt"
	"testing"
	"encoding/json"
	"os"
)

func TestHoge(t *testing.T) {
	s := NewSecret()
	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println(err)
	}
}
