package greeting

import (
	"regexp"
	"testing"
)

// testing Hello function for Valid Name
func TestHelloWithValidName(t *testing.T) {
	name := "bhavya"
	// want := fmt.Sprintf(" Hello, %v Welcome!", name)
	// We have dyamaic message format so we can't predict the exact message
	want := regexp.MustCompile(`\b` + name + `\b`)
	got, err := Hello(name)

	if !want.MatchString(got) || err != nil {
		t.Errorf(`Hello("bhavya") = %q, %v, want match for %#q , nil`, got, err, want)
	}
}

// testing Hello function for Valid Name
func TestHelloWithEmptyName(t *testing.T) {
	name := ""
	want := ""
	got, err := Hello(name)

	if got != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "%v" , error %v`, got, err, want, err)
	}
}
