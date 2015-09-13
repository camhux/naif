package main

import (
	"fmt"
	"strings"
)

func capitalize(s string) string {
	slice := strings.Split(s, "")
	slice[0] = strings.ToUpper(slice[0])
	return strings.Join(slice, "")
}

func makeFileName(forkName, version string) string {
	return fmt.Sprint(capitalize(forkName), " ", version, "-naif", ".sublime-build")
}
