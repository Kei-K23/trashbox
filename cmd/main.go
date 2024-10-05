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
	// time.Sleep(time.Second * 1)
	// err = trashbox.MoveToTrash("test1.txt")
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// }

	files, err := trashbox.ListTrash()
	if err != nil {
		fmt.Println("HERE")
	}

	for _, v := range files {
		fmt.Printf("%s %d\n", v.OriginalPath, v.Size)
	}
	// trashbox.PutBackFromTrash("test.txt")
	// trashbox.PutBackFromTrash("test1.txt")
	trashbox.DeletePermanently("test.txt")

	files, err = trashbox.ListTrash()
	if err != nil {
		fmt.Println("HERE")
	}

	for _, v := range files {
		fmt.Printf("%s %d\n", v.OriginalPath, v.Size)
	}
}
