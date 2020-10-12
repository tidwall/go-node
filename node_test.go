// Copyright 2020 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package node

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	ch := make(chan bool)
	vm := New(&Options{
		OnEmit: func(arg string) {
			if arg != "100" {
				t.Fatalf("expected '%v', got '%v'", "100", arg)
			} else {
				ch <- true
			}
		},
	})
	vm.Run("x=0")
	N := 100
	for i := 0; i < N; i++ {
		val := vm.Run("++x")
		if err := val.Error(); err != nil {
			t.Fatal(err)
		}
		if val.String() != fmt.Sprint(i+1) {
			t.Fatalf("expected '%v', got '%v'", fmt.Sprint(i+1), val)
		}
	}
	val := vm.Run("x")
	if err := val.Error(); err != nil {
		t.Fatal(err)
	}
	if val.String() != fmt.Sprint(N) {
		t.Fatalf("expected '%v', got '%v'", fmt.Sprint(N), val)
	}
	vm.Run("emit(100)")
	<-ch

	v := vm.Run("throw new Error('hello')")
	if v.Error() == nil {
		t.Fatal("expected an error")
	}
	err, ok := v.Error().(ErrThrown)
	if !ok {
		t.Fatal("expected an ErrThrown")
	}
	if err.Error() != "Error: hello" {
		t.Fatalf("expected '%s', got '%s'", "Error: hello", err.Error())
	}

}
