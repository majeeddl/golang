package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	/*
		Create a new sub-directory in the current working directory.
	*/
	err := os.Mkdir("subdir", 0755) // 0755 is the permission mode
	check(err)

	/*
		When creating temporary directories, it’s good practice to defer their removal.
		os.RemoveAll will delete a whole directory tree (similarly to rm -rf).
	*/
	defer os.RemoveAll("subdir") // clean up

	/*
		Helper function to create a new empty file.
	*/
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1") // create a file

	/*
		We can create a hierarchy of directories, including parents with MkdirAll. This is similar to the command-line mkdir -p.
	*/
	err = os.MkdirAll("subdir/parent/child", 0755) // create a hierarchy of directories
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	/*
		ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	*/
	c, err := os.ReadDir("subdir/parent") // list directory contents
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	/*
		Chdir lets us change the current working directory, similarly to cd.
	*/
	err = os.Chdir("subdir/parent/child") // change the current working directory
	check(err)

	/*
		Now we’ll see the contents of subdir/parent/child when listing the current directory.
	*/
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	/*
		cd back to where we started.\
	*/
	err = os.Chdir("../../..")
	check(err)

	/*
		We can also visit a directory recursively, including all its sub-directories. WalkDir accepts a callback function to handle every file or directory visited.
	*/
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

/*
visit is called for every file or directory found recursively by filepath.WalkDir.
*/
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
