package test

import "LINTCODE-GO/src/solutions"
import "testing"

func TestHelloWorld(t *testing.T) {
	if solutions.GetHelloWorld() != "Hello world!" {
		t.Fatalf("Can't get hello world!")
	}
}
