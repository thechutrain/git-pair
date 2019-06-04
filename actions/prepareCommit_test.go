package actions

import (
	"reflect"
	"testing"
)

func TestContainsCoAuthors(t *testing.T) {
	lines := []string{"This does not", "contain Co-authored-by"}
	if containsCoAuthor(lines) {
		t.Error("Lines without the co-author tag were reported to contain it")
	}

	lines2 := []string{"This does contain", "Co-authored-by: cookie monster"}
	if !containsCoAuthor(lines2) {
		t.Error("Lines with the co-author tag were not reported to contain it")
	}

	lines3 := []string{"This does contain", "Co-AuthoRED-By: cookie monster"}
	if !containsCoAuthor(lines3) {
		t.Error("Lines with the co-author tag were not reported to contain it")
	}
}

func TestDeepEqual(t *testing.T) {
	lines := []string{"A", "B", "C"}
	lines2 := []string{"A", "B", "D"}
	lines3 := []string{"A", "B", "C"}

	if reflect.DeepEqual(lines, lines2) {
		t.Error("Unequal lines were reported to be deepequal")
	}

	if !reflect.DeepEqual(lines, lines3) {
		t.Error("Equal lines were reported to be notdeepequal")
	}
}

func TestAddCoAuthors(t *testing.T) {
	lines := []string{"A", "B", "# C"}
	coauthors := []string{"D", "E"}
	expectedResult := []string{"A", "B", "# Added by 🐙", "D", "E", "", "# C"}
	actualResult := addCoAuthors(lines, coauthors)
	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Error("Coauthors were not inserted correctly")
	}
}
