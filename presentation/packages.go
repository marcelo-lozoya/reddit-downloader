package presentation

import (
	"fmt"

	"github.com/marcelo-lozoya/reddit-downloader/presentation/constants"
)

func DoSomethingExported() {
	fmt.Println(constants.ExportedOne) // prints ExportedOne to stdout
}

func doSomethingUnexported() {
	fmt.Println(constants.ExportedOne) // prints ExportedOne to stdout
}
