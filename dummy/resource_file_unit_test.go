package dummy

import (
	"testing"
)

func TestFormatContentsLineOk(t *testing.T) {

	good := "Good contents"
	expected := "CONTENTS: " + good
	l, err := formatContentsLine(good)
	if err != nil {
		t.Errorf("'%s' should not return error", good)
	}
	if l != expected {
		t.Errorf("Expected '%s', got '%s'", expected, l)
	}

	bad := ""
	_, err = formatContentsLine(bad)
	if err == nil {
		t.Errorf("'%s' should return error", bad)
	}
}

// CMG: REMOVE THIS
func savePointer(p *string, v string) {
	*p = v
}

func TestStringPointer(t *testing.T) {
	one := ""
	two := ""

	savePointer(&one, "one")
	savePointer(&two, "two")

	if one != "one" {
		t.Errorf("'%s' is not 'one'", one)
	}
	if two != "two" {
		t.Errorf("'%s' is not 'two'", two)
	}
}
