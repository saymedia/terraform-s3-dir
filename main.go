package main

import (
	"fmt"
	"os"

	getopt "code.google.com/p/getopt"
)

var exclude = getopt.ListLong("exclude", 'x', "", "glob patterns to exclude")
var help = getopt.BoolLong("help", 'h', "", "print this help")

func main() {
	getopt.SetParameters("<root dir> <bucket name>");
	getopt.Parse()
	if (*help) {
		getopt.PrintUsage(os.Stdout)
		return
	}

	args := getopt.Args()
	if len(args) != 2 {
		getopt.PrintUsage(os.Stderr)
		os.Exit(1)
	}

	rootDir := args[0]
	bucketName := args[1]

	fmt.Println(rootDir, bucketName)

}
