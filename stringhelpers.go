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

func makeVariantName(fork, version string) string {
	if fork == "node" {
		fork = capitalize(fork)
	}
	return fmt.Sprint(fork, " ", version)
}
