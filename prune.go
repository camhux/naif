package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func checkVerFromBuild(fileName string, builds []BuildTemplate) bool {
	for _, build := range builds {
		if build.filename == fileName {
			return true
		}
	}
	return false
}

func pruneSavedBuilds(sublimepath string, builds []BuildTemplate) {
	r := regexp.MustCompile("^(Node|io.js) v\\d+\\.\\d+\\.\\d-naif.sublime-build$")

	sublimeDir, err := os.Open(sublimepath)
	if err != nil {
		log.Fatal(err)
	}
	defer sublimeDir.Close()

	filenames, err := sublimeDir.Readdirnames(-1)

	var buildNamesFromFile []string

	for _, filename := range filenames {
		if r.MatchString(filename) {
			buildNamesFromFile = append(buildNamesFromFile, filename)
		}
	}

	for _, buildName := range buildNamesFromFile {
		if !checkVerFromBuild(buildName, builds) {
			os.Remove(filepath.Join(sublimepath, buildName))
			log.Printf("Removed build %v", buildName)
		}
	}
}
