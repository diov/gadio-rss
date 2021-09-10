package main

import "testing"

func TestEndPoint(t *testing.T) {
	endpoint := generateEndPoint(10)
	t.Log(endpoint)
}

func TestFetch(t *testing.T) {
	fetch(0, false)
}
