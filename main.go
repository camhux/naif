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

	path := filepath.Join(nvmpath, "versions", fork, version, "bin")

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
	setSublimePath()

	var ok bool // separate declaration required, since multiple short assignment shadows the global nvmpath var
	nvmpath, ok = os.LookupEnv("NVM_DIR")
	if !ok {
		log.Fatal("Unable to locate .nvm directory!")
	}

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

func checkOrWriteBuild(sublimepath string, build BuildTemplate) {
	writepath := filepath.Join(sublimepath, build.filename)

	json, err := json.Marshal(build)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(writepath); err != nil {
		writeErr := ioutil.WriteFile(writepath, json, 0644)
		if writeErr != nil {
			log.Fatal(writeErr)
		}
		log.Printf("Wrote %v in %v ", build.filename, sublimepath)
	} else {
		log.Printf("Build system %v already exists, leaving in place", build.filename)
	}
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
	defer forkDir.Close()

	forkVers, err := forkDir.Readdirnames(-1)
	if err != nil {
		log.Fatal(err)
	}

	return forkVers
}

func setSublimePath() {
	parent := filepath.Join(homepath, "Library", "Application Support")
	st2 := "Sublime Text 2"
	st3 := "Sublime Text 3"
	end := filepath.Join("Packages", "User")

	if _, err := os.Stat(filepath.Join(parent, st3)); err == nil {
		sublimepath = filepath.Join(parent, st3, end)
	} else if _, err := os.Stat(filepath.Join(parent, st2)); err == nil {
		sublimepath = filepath.Join(parent, st2, end)
	} else {
		log.Fatal("Cannot find SublimeText directory!")
	}
}
