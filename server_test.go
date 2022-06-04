package main

import (
	"io/ioutil"
	"os"
	"testing"

	postgrescontainer "github.com/GolangStudiy/go-users-postgres-rest-api/postgrescontainer"
)

func TestShouldBePrintCorrectMessage(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	postgrescontainer.NewTestDatabase(t)
	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != "App Started" {
		t.Errorf("Expected %s, got %s", "App Started", out)
	}
}
