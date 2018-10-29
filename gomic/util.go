package gomic

import (
	"log"
	"testing"
)

type (
	// CallbackNotImplemented is a type of mock's CallbackNotImplemented field.
	CallbackNotImplemented func(t *testing.T, intfName, methodName string)
)

// DefaultCallbackNotImplemented is a default function which is called when the method is not implemented.
func DefaultCallbackNotImplemented(t *testing.T, intf, method string) {
	if t == nil {
		log.Fatalf("%s#%s is not implemented", intf, method)
		return
	}
	t.Fatalf("%s#%s is not implemented", intf, method)
}

// DoNothing does nothing.
func DoNothing(t *testing.T, intf, method string) {}
