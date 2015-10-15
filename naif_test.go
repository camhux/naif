package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

var vs = Variants{
	NewVariant("node", "v0.12.7"),
	NewVariant("node", "v4.0.0"),
	NewVariant("io.js", "v3.2.1"),
}

func TestVariantsLess(t *testing.T) {
	tvs := make(Variants, len(vs))
	copy(tvs, vs)

	if isLess := tvs.Less(0, 1); !isLess {
		t.Errorf("Expected %v to be judged less than %v", tvs[0], tvs[1])
		t.Log(tvs)
	}

	if isLess := tvs.Less(1, 2); isLess {
		t.Errorf("Expected %v to be judged greater than %v", tvs[1], tvs[2])
		t.Log(tvs)
	}
}

func TestVariantsSwap(t *testing.T) {
	tvs := make(Variants, len(vs))
	copy(tvs, vs)

	tvs.Swap(0, 2)
	if !strings.Contains(tvs[0].Name, "io.js") {
		t.Errorf("Expected %v to have name \"io.js\"", tvs[0].Name)
		t.Log(tvs)
	}
	if !strings.Contains(tvs[1].Name, "v4.0.0") {
		t.Errorf("Expected %v to have version \"v4.0.0\"", tvs[1].Name)
		t.Log(tvs)
	}
	if !strings.Contains(tvs[2].Name, "v0.12.7") {
		t.Errorf("Expected %v to have version \"v0.12.7\"", tvs[2].Name)
		t.Log(tvs)
	}
}

func TestVariantsSort(t *testing.T) {
	tvs := make(Variants, len(vs))
	copy(tvs, vs)

	sort.Sort(tvs)

	if actual, expected := tvs[0].Name, "Node v0.12.7"; actual != expected {
		t.Errorf("Expected %v to be %v", actual, expected)
		t.Log(tvs)
	}
	if actual, expected := tvs[1].Name, "io.js v3.2.1"; actual != expected {
		t.Errorf("Expected %v to be %v", actual, expected)
		t.Log(tvs)
	}
	if actual, expected := tvs[2].Name, "Node v4.0.0"; actual != expected {
		t.Errorf("Expected %v to be %v", actual, expected)
		t.Log(tvs)
	}
}

func TestBuildTemplate(t *testing.T) {
	tvs := make(Variants, len(vs))
	copy(tvs, vs)

	b, err := NewBuildTemplate(tvs)
	if err != nil {
		t.Error(err)
	}

	fmt.Print(b)

	sort.Sort(tvs)

	if actual, expected := b.Cmd, tvs[len(tvs)-1].Cmd; !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v to be %v", actual, expected)
	}

	if actual, expected := b.Path, tvs[len(tvs)-1].Path; actual != expected {
		t.Errorf("Expected %v to be %v", actual, expected)
	}

	bytes, err := json.Marshal(b)

	fmt.Println(string(bytes))

}
