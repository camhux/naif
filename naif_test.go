package main

import (
	"testing"
)

func TestVariantsLess(t *testing.T) {
	vs := Variants{
		NewVariant("node", "v0.12.7"),
		NewVariant("node", "v4.0.0"),
		NewVariant("io.js", "v3.2.1"),
	}

	if isLess := vs.Less(0, 1); !isLess {
		t.Errorf("Expected %v to be judged less than %v", vs[0], vs[1])
	}

	if isLess := vs.Less(1, 2); isLess {
		t.Errorf("Expected %v to be judged greater than %v", vs[1], vs[2])
	}
}
