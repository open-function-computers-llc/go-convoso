package convoso

import "testing"

func TestCanNotInitializeThePackageWithoutAnAPIKey(t *testing.T) {
	c := Config{}
	err := Init(c)

	if err == nil {
		t.Errorf("You shouldn't be able to initialize this package without setting an API key in the configuration")
	}

	c.APIKey = "anythingCanGoHere"
	err = Init(c)

	if err != nil {
		t.Errorf("The Init func should work with the valid config struct")
	}
}
