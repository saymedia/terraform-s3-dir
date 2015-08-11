package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	getopt "code.google.com/p/getopt"
)

var exclude = getopt.ListLong("exclude", 'x', "", "glob patterns to exclude")
var help = getopt.BoolLong("help", 'h', "", "print this help")

func main() {
	getopt.SetParameters("<root dir> <bucket name>")
	getopt.Parse()
	if *help {
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

	resourcesMap := map[string]interface{}{}
	result := map[string]interface{}{
		"resource": map[string]interface{}{
			"aws_s3_bucket_object": resourcesMap,
		},
	}

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %s\n", path, err)
			// Skip stuff we can't read.
			return nil
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed make %s relative: %s\n", path, err)
			return nil
		}

		path, err = filepath.EvalSymlinks(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to resolve symlink %s: %s\n", path, err)
			return nil
		}

		if info.IsDir() {
			// Don't need to create directories since they are implied
			// by the files within.
			return nil
		}

		for _, pattern := range *exclude {
			var toMatch []string
			if strings.ContainsRune(pattern, filepath.Separator) {
				toMatch = append(toMatch, relPath)
			} else {
				// If the pattern does not include a path separator
				// then we apply it to all segments of the path
				// individually.
				toMatch = strings.Split(relPath, string(filepath.Separator))
			}

			for _, matchPath := range toMatch {
				matched, _ := filepath.Match(pattern, matchPath)
				if matched {
					return nil
				}
			}
		}

		resourceNameBytes := sha1.Sum([]byte(relPath))
		resourceName := fmt.Sprintf("%x", resourceNameBytes)

		resourcesMap[resourceName] = map[string]interface{}{
			"bucket": bucketName,
			"key":    relPath,
			"source": path,
		}

		return nil
	})

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(result)
}
