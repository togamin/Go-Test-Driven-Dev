package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("togamin")
    want := "Hello, togamin"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}