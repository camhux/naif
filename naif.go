package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type SublBuildTemplate struct {
}

// func (s *SublBuildTemplate) MarshallJSON() ([]byte, error) {

// }

func main() {

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

	var verPaths []string
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
			verPaths = append(verPaths, filepath.Join(forkDir.Name(), version, "bin"))
		}
	}
	fmt.Print(verPaths)

}
