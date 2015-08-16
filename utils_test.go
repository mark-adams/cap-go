package cap

import (
	"strings"
	"testing"

	"github.com/kr/pretty"
)

func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Error(message)
		for _, desc := range pretty.Diff(expected, actual) {
			t.Error(desc)
		}
	}
}

func assertStartsWith(t *testing.T, strVal, prefix string, message string) {

	if strings.Index(strVal, prefix) == -1 {
		t.Error(message)
		for _, desc := range pretty.Diff(prefix, strVal[:len(prefix)]) {
			t.Error(desc)
		}
	}
}

func assertIn(t *testing.T, needle string, haystack []string, message string) {
	found := false

	for _, value := range haystack {
		if value == needle {
			found = true
			break
		}
	}

	if !found {
		t.Error(message)
		t.Errorf("Needle '%s' not found in '%s'", needle, haystack)
	}
}
