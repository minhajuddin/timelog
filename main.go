package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) == 1 {
		printLog(1)
	}
}

func printLog(days int) {
	f, err := os.Open("/home/minhajuddin/.timelog")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	io.Copy(os.Stdout, f)
}
