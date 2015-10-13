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
	homepath    = getHomePath()
	sublimepath = getSublimePath()
	nvmpath     = os.Getenv("NVM_DIR")
)

type BuildTemplate struct {
	Cmd      [2]string `json:"cmd"`
	Path     string    `json:"path"`
	Selector string    `json:"selector"`
	Variants []Variant `json:"variants,omitempty"`
	filename string
}

func NewBuildTemplate(variants []Variant) BuildTemplate {
	if len(variants) == 0 {
		log.Fatal("No build systems to write. Exiting...")
	}
	defaultVar := variants[0]
	restVariants := variants[1:]

	return BuildTemplate{
		Cmd:      [2]string{cmd, "$file"},
		Path:     path,
		Selector: "source.js",
		Variants: restVariants,
		filename: "Node (naif)",
	}
}

type Variant struct {
	Name string    `json:"name"`
	Path string    `json:"path"`
	Cmd  [2]string `json:"cmd,omitempty"`
}

func newVariant(fork, version string) Variant {
	name := makeVariantName(fork, version)
	path := filepath.Join(nvmpath, "versions", fork, version, "bin")
	cmd := strings.Replace(fork, ".", "", 1)

	return Variant{
		name,
		path,
		[2]string{cmd, "$file"},
	}
}

func main() {

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

func getSublimePath() string {
	var s string
	parent := filepath.Join(homepath, "Library", "Application Support")
	st2 := "Sublime Text 2"
	st3 := "Sublime Text 3"
	end := filepath.Join("Packages", "User")

	if _, err := os.Stat(filepath.Join(parent, st3)); err == nil {
		s = filepath.Join(parent, st3, end)
	} else if _, err := os.Stat(filepath.Join(parent, st2)); err == nil {
		s = filepath.Join(parent, st2, end)
	} else {
		log.Fatal("Cannot find SublimeText directory!")
	}

	return s
}

func getHomePath() string {
	currUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return currUser.HomeDir
}
