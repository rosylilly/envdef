package envdef

import (
	"os"
	"testing"
)

func prepare() {
	os.Clearenv()

	os.Setenv("FOUND_ENV", "1")
	os.Setenv("EXPAND_ENV", "expanded ${FOUND_ENV}")
}

func TestGet(t *testing.T) {
	prepare()

	found := Get("FOUND_ENV", "2")

	if found != "1" {
		t.Errorf("Missmatch %s : %s", found, "1")
	}

	notFound := Get("NOT_FOUND_ENV", "2")

	if notFound != "2" {
		t.Errorf("Missmatch %s : %s", notFound, "2")
	}
}

func TestGetWithExpand(t *testing.T) {
	expanded := GetWithExpand("EXPAND_ENV", "not found")

	if expanded != "expanded 1" {
		t.Errorf("Missmatch %s", expanded)
	}

	expandedNotFound := GetWithExpand("NOT_FOUND_EXPAND_ENV", "not found with ${FOUND_ENV}")

	if expandedNotFound != "not found with 1" {
		t.Errorf("Missmatch %s", expandedNotFound)
	}
}
