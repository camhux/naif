package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var (
	homepath    string
	sublimepath string
	nvmpath     string
)

type BuildTemplate struct {
	Cmd      [2]string `json:"cmd"`
	Path     string    `json:"path"`
	filename string
}

func NewBuildTemplate(fork, version string) BuildTemplate {
	var cmd string
	if strings.Contains(fork, ".") {
		cmd = strings.Replace(fork, ".", "", 1)
	} else {
		cmd = fork
	}

	path := filepath.Join(nvmpath, fork, version, "bin")

	return BuildTemplate{
		Cmd:      [2]string{cmd, "$file"},
		Path:     path,
		filename: makeFileName(fork, version),
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	homepath = user.HomeDir
	sublimepath = filepath.Join(homepath, "Library/Application Support/Sublime Text 3/Packages/User")

	nvmpath, ok := os.LookupEnv("NVM_DIR")
	if !ok {
		log.Fatal("Unable to locate .nvm directory!")
	}
	// aliasDir := filepath.Join(nvmpath, "alias")

	var builds []BuildTemplate

	forks := getForknames()
	for _, fork := range forks {
		versions := getVersOfFork(fork)
		for _, version := range versions {
			builds = append(builds, NewBuildTemplate(fork, version))
		}
	}

	for _, build := range builds {
		checkOrWriteBuild(sublimepath, build)
	}
	pruneSavedBuilds(sublimepath, builds)
}

func checkOrWriteBuild(sublimepath string, build BuildTemplate) error {
	writepath := filepath.join(sublimepath, build.filename)

	json, err := json.Marshal(build)
	if err != nil {
		return err
	}

	if _, err := os.Stat(writepath); err != nil {
		writeErr := ioutil.WriteFile(writepath, json, 0644)
		if writeErr != nil {
			return writeErr
		}
	}
	return nil
}

func getForknames() []string {
	versionspath := filepath.Join(nvmpath, "versions")

	versionsDir, err := os.Open(versionspath)
	if err != nil {
		log.Fatalf("Unable to read versions directory: %v", err)
	}
	defer versionsDir.Close()

	forknames, err := versionsDir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	return forknames
}

func getVersOfFork(forkname string) []string {
	forkpath := filepath.Join(nvmpath, "versions", forkname)

	forkDir, err := os.Open(forkpath)
	if err != nil {
		log.Fatalf("Unable to read fork directory: %v", err)
	}
	defer forkDir.close()

	forkVers, err := forkDir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	return forkVers
}
