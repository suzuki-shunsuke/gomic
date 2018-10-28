package gomic

import (
	"log"
	"testing"
)

// DefaultCallbackNotImplemented is a default function which is called when the method is not implemented.
func DefaultCallbackNotImplemented(t *testing.T, intf, method string) {
	if t == nil {
		log.Fatalf("%s#%s is not implemented", intf, method)
		return
	}
	t.Fatalf("%s#%s is not implemented", intf, method)
}
