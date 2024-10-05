package main

import (
	"fmt"

	"github.com/Kei-K23/trashbox"
)

func main() {
	// err := trashbox.MoveToTrash("test.txt")
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }

	files, err := trashbox.ListTrash()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, file := range files {
		fmt.Printf("%s %d %s", file.OriginalPath, file.Size, file.DeletedAt.String())
	}

	err = trashbox.PutBackFromTrash("test.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	files, err = trashbox.ListTrash()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, file := range files {
		fmt.Printf("%s %d %s", file.OriginalPath, file.Size, file.DeletedAt.String())
	}

}
