package main

import "testing"

/*
* $ go test -v
 */

func TestSum(t *testing.T) {
	// GetSlug : testing this
	slug := GetSlug("test@gmail.com")
	neededResult := "gmail.com"
	if slug != neededResult {
		t.Errorf("Slug was incorrect, got: %s, want: %s.", slug, neededResult)
	}
}
