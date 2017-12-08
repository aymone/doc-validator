// +build acceptance

package main

import (
	"os"
	"testing"
)

const host string = "http://localhost:3000"

func TestMain(m *testing.M) {
	//Drop all collection before run tests
	getClient().C("documents").DropCollection()

	// call tests
	os.Exit(m.Run())
}
