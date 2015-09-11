package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	homeDir := user.HomeDir
	sublimepath := filepath.Join(homeDir, "Library/Application Support/Sublime Text 3/Packages/User")

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
		}
	}

	for _, build := range builds {
		writePath := filepath.Join(sublimepath, build.filename)
		json, err := json.Marshal(build)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(writePath, json, 0644)
	}
	pruneSavedBuilds(sublimepath, builds)
}
