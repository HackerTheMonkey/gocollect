package gocollect

import "testing"

func TestFilteringProducesTheExpectedOutcome(t *testing.T) {
	// Given
	expectedArrayLength := 1
	expectedFilteredEntry := "Mars"

	planets := []interface{}{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
		"Pluto",
	}

	isMarsPredicate := func(value interface{}) bool {
		return value.(string) == "Mars"
	}

	// When
	underlyingArray := CollectionFrom(planets).Filter(isMarsPredicate).Unwrap()

	// Then
	arrayLength := len(underlyingArray)
	if arrayLength != expectedArrayLength{
		t.Errorf("Expecting a single element in the array but found %d", arrayLength)
	}

	actualArrayElement := underlyingArray[0]
	if actualArrayElement != expectedFilteredEntry{
		t.Errorf("Expected: %s, Actual: %s", expectedFilteredEntry, actualArrayElement)
	}
}
