package main

import "strings"

type BuildTemplate struct {
	Cmd      [2]string `json:"cmd"`
	Path     string    `json:"path"`
	filename string
}

func NewBuildTemplate(path, fork, version string) BuildTemplate {
	var cmd string
	if strings.Contains(fork, ".") {
		cmd = strings.Replace(fork, ".", "", 1)
	} else {
		cmd = fork
	}

	return BuildTemplate{
		Cmd:      [2]string{cmd, "$file"},
		Path:     path,
		filename: makeFileName(fork, version),
	}
}
