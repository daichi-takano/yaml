package main

import (
	"os"
	"fmt"
)

func f_out(data string) {
	// create new file
	new_file, err := os.OpenFile("create_new_file.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error")
		return
	}

	// write in the file
	defer new_file.Close()
	// fmt.Fprintln(filename, contents

	fmt.Fprintf(new_file, data)
}