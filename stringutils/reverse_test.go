
package stringutil

import "testing"

func TestReversesString(t *testing.T) {

	reversed := Reverse("Testing")
	if reversed != "gnitseT" {
		t.Errorf("Got %q instead of 'gnitseT'", reversed)
	}

}

func TestReverseHandlesEmptyString(t *testing.T){

	reversed := Reverse("")
	if reversed != "" {
		t.Errorf("Reversing empty string failed, got %q", reversed)
	}

}
