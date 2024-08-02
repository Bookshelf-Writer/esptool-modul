package parser

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	maps, err := LoadIndex()
	if err != nil {
		t.Fatal(err)
	}

	for key, val := range *maps {
		fmt.Println(key, val)
	}
}
