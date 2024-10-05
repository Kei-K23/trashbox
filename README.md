# Trashbox 🗑️

[![Go Test](https://github.com/Kei-K23/trashbox/actions/workflows/test.yml/badge.svg)](https://github.com/Kei-K23/trashbox/actions/workflows/test.yml)

**Trashbox** 🗑️ is a small, fast and cross-platform **Go** utility package for dealing with OS native `Trash Box` or `Recycle Bin` according to user OS. This package allows you to move files to the trash, restore them, list items in the trash, delete items permanently, and automatically clean up old items based on a specified duration.

⚠️❗️
Trashbox is currently under development and only tested in `Mac OS`. And also, intensive changes can happen to the package at any time.
❗️⚠️

## Table of Contents

- [Trashbox 🗑️](#trashbox-️)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Move to Trash](#move-to-trash)
    - [Put Back from Trash](#put-back-from-trash)
    - [List Trash](#list-trash)
    - [Delete Permanently](#delete-permanently)
    - [Auto Clean Trash](#auto-clean-trash)
  - [License](#license)
  - [Contribution](#contribution)

## Installation

To use the **Trashbox** package, first ensure you have Go installed on your machine. You can then get the package using the following command:

```bash
go get github.com/Kei-K23/trashbox
```

## Usage

### Move to Trash

To move a file to the trash, use the `MoveToTrash` function. This function takes the file path as an argument.

```go
package main

import (
	"fmt"
	"log"

	"github.com/Kei-K23/trashbox"
)

func main() {
	err := trashbox.MoveToTrash("path/to/your/file.txt")
	if err != nil {
		log.Fatalf("Error moving to trash: %v", err)
	}
	fmt.Println("File moved to trash successfully.")
}
```

### Put Back from Trash

To restore a file from the trash, use the `PutBackFromTrash` function. Pass the filename (not the full path) that you want to restore.

```go
err := trashbox.PutBackFromTrash("file.txt")
if err != nil {
	log.Fatalf("Error putting back from trash: %v", err)
}
fmt.Println("File restored from trash successfully.")
```

### List Trash

To get a list of files in the trash, use the `ListTrash` function. This function returns a slice of metadata containing details about each file.

```go
files, err := trashbox.ListTrash()
if err != nil {
	log.Fatalf("Error listing trash: %v", err)
}

for _, file := range files {
	fmt.Printf("File: %s, Deleted At: %s, Size: %d bytes\n", file.Filename, file.DeletedAt, file.Size)
}
```

### Delete Permanently

To delete a file permanently from the trash, use the `DeletePermanently` function. Pass the filename (not the full path) of the file you want to delete.

```go
err := trashbox.DeletePermanently("file.txt")
if err != nil {
	log.Fatalf("Error deleting permanently: %v", err)
}
fmt.Println("File deleted permanently.")
```

### Auto Clean Trash

To automatically clean up files older than a specified number of days, use the `AutoCleanTrash` function. You can pass one or zero arguments, where the argument is the number of days.

```go
err := trashbox.AutoCleanTrash(30) // Deletes files older than 30 days
if err != nil {
	log.Fatalf("Error cleaning trash: %v", err)
}
fmt.Println("Old files cleaned from trash.")
```

## License

This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.

## Contribution

All contributions are very welcome. Please open issues or make PR for error, bug, and adding new features. If you have any additional features or changes in the package, consider updating the documentation accordingly! Let me know if you need further modifications or additional information!
