package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func examUnzip() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("testdata/edwin.txt.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {

		fmt.Printf("Contents of %s: %T\n", f.Name, *f)
		fmt.Printf(f.FileHeader.FileInfo().ModTime().String())
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, rc)
		println()

		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println("-----------------------------------")
	}
}
