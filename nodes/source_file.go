// an input node
// reads a single text file

// 2019-04-20 owi - this is my first LogFlow node

package source_file

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(b))
}
