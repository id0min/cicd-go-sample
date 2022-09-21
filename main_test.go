package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello World!!1" {
		t.Fatal("Test fail")
	}
}
