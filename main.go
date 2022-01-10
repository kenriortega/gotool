package main

import (
	"log"
	"os"
)

var (
	folders = []string{
		"basic/cmd",
		"basic/pkg",
		"basic/internal/domain",
		"basic/internal/ports",
		"basic/internal/handlers",
		"basic/internal/services",
	}
)

func main() {
	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
