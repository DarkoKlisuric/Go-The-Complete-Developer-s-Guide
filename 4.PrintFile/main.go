package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	// list of argument that have been sent to program
	//example: go run main.go myFile.txt
	// fmt.Println(os.Args[0]) will print -> /tmp/go-build344630151/b001/exe/main
	// fmt.Println(os.Args[1]) will print myFile.txt

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f) 
}
