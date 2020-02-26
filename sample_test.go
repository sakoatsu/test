package main

import "testing"

func TestSample(t *testing.T) {
    res := hoge("aaa")
    if res != "bbb" {
        t.Fatalf("failed test")
    }
}
