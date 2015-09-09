package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

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

func capitalize(s string) string {
	slice := strings.Split(s, "")
	slice[0] = strings.ToUpper(slice[0])
	return strings.Join(slice, "")
}

func makeFileName(forkName, version string) string {
	return fmt.Sprint(capitalize(forkName), " ", version, ".sublime-build")
}

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	homeDir := user.HomeDir
	STDir := filepath.Join(homeDir, "Library/Application Support/Sublime Text 3/Packages/User")

	nvmDir, ok := os.LookupEnv("NVM_DIR")
	if !ok {
		log.Fatal("Unable to locate .nvm directory!")
	}
	// aliasDir := filepath.Join(nvmDir, "alias")

	versionsDir := filepath.Join(nvmDir, "versions")

	versions, err := os.Open(versionsDir)
	if err != nil {
		log.Fatalf("Unable to read versions directory: %v", err)
	}
	defer versions.Close()

	forkNames, err := versions.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	var builds []BuildTemplate
	for _, fork := range forkNames {

		forkDir, err := os.Open(filepath.Join(versionsDir, fork))
		if err != nil {
			log.Fatalf("Unable to open fork subdirectory: %v Err: ", fork, err)
		}
		defer forkDir.Close()

		forkVersionNames, err := forkDir.Readdirnames(-1)
		if err != nil {
			log.Fatal(err)
		}

		for _, version := range forkVersionNames {
			path := filepath.Join(forkDir.Name(), version, "bin")
			builds = append(builds, NewBuildTemplate(path, fork, version))
			// verPaths = append(verPaths, filepath.Join(forkDir.Name(), version, "bin"))
		}
	}

	for _, build := range builds {
		writePath := filepath.Join(STDir, build.filename)
		json, err := json.Marshal(build)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Print(writePath, string(json))
		ioutil.WriteFile(writePath, json, 0644)
	}
}
