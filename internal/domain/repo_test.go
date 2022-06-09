package domain

import "testing"

func TestIsValid(t *testing.T) {
	repo := Repository{}
	if repo.isValid() {
		t.Fatalf("emtpy repo is valid")
	}
}
