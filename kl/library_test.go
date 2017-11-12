package kl

import (
	// "fmt"
	"testing"
)

func TestReverse1(t *testing.T) {
	if reverse1(Nil) != Nil {
		t.FailNow()
	}

	l := cons(Make_integer(1), cons(Make_integer(2), cons(Make_integer(3), Nil)))
	r := reverse1(l)
	if mustInteger(car(r)) != 3 {
		t.FailNow()
	}
	if mustInteger(cadr(r)) != 2 {
		t.FailNow()
	}
	if mustInteger(caddr(r)) != 1 {
		t.FailNow()
	}
	if cdddr(r) != Nil {
		t.FailNow()
	}

	// (1 (1 2 3))
	l1 := cons(Make_integer(1), cons(l, Nil))
	// ((1 2 3) 1)
	l2 := cons(l, cons(Make_integer(1), Nil))

	// fmt.Printf("%#v...\n", (*scmHead)(l1))
	// fmt.Printf("%#v...\n", (*scmHead)(l2))

	if equal(reverse(l1), l2) != True {
		t.Error("fuck1")
	}
	if equal(reverse(l2), l1) != True {
		t.Error("fuck2")
	}
}
