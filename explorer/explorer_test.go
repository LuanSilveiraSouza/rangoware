package explorer

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestList(t *testing.T) {
	files := MapFiles()

	for _, v := range files {
		file, err := ioutil.ReadFile(v)

		fmt.Println(v)

		if err != nil {
			continue
		}

		fmt.Println(file)
		fmt.Println(string(file))
	}

	if len(files) > 0 {
		t.Fatal("MapFiles not working as expected")
	}
}
