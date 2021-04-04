package explorer

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	files := MapFiles()

	for _, v := range files {
		fmt.Println(v)
	}

	if len(files) > 0 {
		t.Fatal("MapFiles not working as expected")
	}
}
