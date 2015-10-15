package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	homepath = getHomePath()
	nvmpath  = getNvmPath()
	destpath = getDestPath()
)

func main() {

	var variants Variants

	forks := getForknames()
	for _, fork := range forks {
		versions := getVersOfFork(fork)
		for _, version := range versions {
			variants = append(variants, NewVariant(fork, version))
		}
	}

	buildTemplate, err := NewBuildTemplate(variants)
	if err != nil {
		log.Fatal(err)
	}

	buildJson, err := json.MarshalIndent(buildTemplate, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	dest, err := os.Create(filepath.Join(destpath, buildTemplate.filename))
	if err != nil {
		log.Fatal(err)
	}

	defer dest.Close()

	log.Println("naif is writing to:", destpath)
	_, err = dest.Write(buildJson)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Done!")
}

type BuildTemplate struct {
	Cmd      []string  `json:"cmd"`
	Path     string    `json:"path"`
	Selector string    `json:"selector"`
	Variants []Variant `json:"variants,omitempty"`
	filename string
}

func NewBuildTemplate(variants Variants) (BuildTemplate, error) {
	if len(variants) == 0 {
		return BuildTemplate{}, errors.New("No builds to write")
	}

	vs := make(Variants, len(variants))
	copy(vs, variants)

	sort.Sort(vs)

	defaultVar := vs[len(vs)-1]

	for i := range vs {
		if vs[i].Cmd[0] == defaultVar.Cmd[0] {
			vs[i].Cmd = make([]string, 0)
		}
	}

	return BuildTemplate{
		Cmd:      defaultVar.Cmd,
		Path:     defaultVar.Path,
		Selector: "source.js",
		Variants: vs,
		filename: "Node (naif).sublime-build",
	}, nil
}

type Variant struct {
	Name string   `json:"name"`
	Path string   `json:"path"`
	Cmd  []string `json:"cmd,omitempty"`
}

func NewVariant(fork, version string) Variant {
	name := makeVariantName(fork, version)
	path := filepath.Join(nvmpath, "versions", fork, version, "bin")
	cmd := strings.Replace(fork, ".", "", 1)

	return Variant{
		name,
		path,
		[]string{cmd, "$file"},
	}
}

type Variants []Variant

var verPattern *regexp.Regexp = regexp.MustCompile("\\d{1,2}")

func (vars Variants) Len() int {
	return len(vars)
}

func (vars Variants) Less(i, j int) bool {
	s := vars
	verA := verPattern.FindAllString(s[i].Name, 3)
	verB := verPattern.FindAllString(s[j].Name, 3)

	for i := range verA {
		segA, errA := strconv.Atoi(verA[i])
		segB, errB := strconv.Atoi(verB[i])

		if errA != nil || errB != nil {
			log.Fatal("Error sorting variants: ", errA, errB)
		}

		if segA < segB {
			return true
		} else {
			break
		}
	}

	return false
}

func (vars Variants) Swap(i, j int) {
	s := vars
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
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

func getHomePath() string {
	currUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return currUser.HomeDir
}

func getNvmPath() string {
	p, ok := os.LookupEnv("NVM_DIR")
	if !ok {
		log.Fatal("$NVM_DIR isn't set. Ensure that nvm is installed and sourced in your shell.")
	}

	return p
}

func getDestPath() string {
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
