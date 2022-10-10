package main

import "testing"

func TestRun(t *testing.T) {
	err := run(1)

	if err != nil {
		t.Error("failed  run ()")
	}
}
