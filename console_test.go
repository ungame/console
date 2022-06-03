package console

import "testing"

func TestLog(t *testing.T) {

	type Any struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	err := Log(Any{"Hello", 10}, " ")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}
