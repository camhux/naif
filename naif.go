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

	versions, err := os.Open(filepath.Join(nvmDir, "versions"))
	if err != nil {
		log.Fatalf("Unable to read versions directory: %v", err)
	}

	versToBuild := versions.Readdirnames(-1)
	verPaths := make([]string, len(versToBuild))

}
