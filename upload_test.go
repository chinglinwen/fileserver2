package main

import (
	"syscall"
	"testing"
)

func TestAppendFlag(t *testing.T) {
	appendmode := "append"
	appendflags := mode(len(appendmode) != 0)
	normalflags := mode(false)
	if appendflags != 1089 || normalflags != 577 {
		t.Errorf("flags setting failed, should not equal, %v = %v", appendflags, normalflags)
		return
	}
	t.Logf("appendflags: %v, normalflags: %v, syscallappend: %v\n", appendflags, normalflags, syscall.O_APPEND)
}
