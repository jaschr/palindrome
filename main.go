/*
 * Copyright (c) 2022 Jacob Schroeder.
 *
 * Palindrome is a (G)o (S)tatic (S)ite (G)enerator.
 * Github Repo: https://github.com/jaschr/palindrome
 * Documentation: https://github.com/jaschr/palindrome/tree/main/docs
 */

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func main() {
	root := "posts/"
	files, err := WalkMatch("./"+root, "*.md")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		mdf := strings.ReplaceAll(f, root, "")
		fmt.Println(mdf)
	}
}
