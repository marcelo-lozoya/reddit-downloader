package golang101

import (
	"fmt"

	"github.com/marcelo-lozoya/reddit-downloader/golang101/constants"
)

func DoSomethingExported() {
	fmt.Println(constants.ExportedOne) // prints ExportedOne to stdout
}

func doSomethingUnexported() {
	fmt.Println(constants.ExportedOne) // prints ExportedOne to stdout
}
