package main

// 1. open the file
// 	1-1 open the txt file (o)
// 	1-2 open the .docx file (--)
// 2. read the contents of the file
// 3. extract and rebuild the contents

import (
	"os"
	"fmt"
)

func extracting() {
	// read the file
	fi, err := os.Open("path_for_file")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	// write the file
	fo, err := os.Create("path_for_file")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	for {
		// reading
		cnt, err := fi.Read(buff)
		if err != nil & err != io.EOF {
			panic(err)
		}

		// loop is ended
		if cnt == 0 {
			break
		}

		// writting
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}