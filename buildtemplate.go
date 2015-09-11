package main

type BuildTemplate struct {
	Cmd      [2]string `json:"cmd"`
	Path     string    `json:"path"`
	filename string
}

func NewBuildTemplate(path, fork, version string) BuildTemplate {
	return BuildTemplate{
		Cmd:      [2]string{fork, "$file"},
		Path:     path,
		filename: makeFileName(fork, version),
	}
}
