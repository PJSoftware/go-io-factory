package factory_test

import (
	"testing"

	factory "github.com/pjsoftware/go-io-factory"
)

func TestQueuePeekEmpty(t *testing.T) {
	queue := factory.NewQueue[int]()

	x := queue.Peek()
	if x != nil {
		t.Errorf("expected nil; got %v", x)
	}
}

func TestQueuePeekNext(t *testing.T) {
	queue := factory.NewQueue[int]()

	queue.Push(1)
	queue.Push(2)

	a := queue.Peek()
	if *a != 1 {
		t.Errorf("peek() expected 1; got %d", *a)
	}

	b := queue.Peek()
	if *b != 1 {
		t.Errorf("peek() expected 1; got %d", *b)
	}

	c := queue.Next()
	if *c != 1 {
		t.Errorf("next() expected 1; got %d", *c)
	}

	d := queue.Next()
	if *d != 2 {
		t.Errorf("next() expected 2; got %d", *d)
	}

	e := queue.Next()
	if e != nil {
		t.Errorf("next() expected nil; got %d", e)
	}

	if b != c {
		t.Errorf("expected b & c pointers to be equal")
	}

	if b == d {
		t.Errorf("expected b & d pointers to NOT be equal")
	}
}

func TestQueuePeekNextPointers(t *testing.T) {
	queue := factory.NewQueue[int]()

	oneA := 1
	oneB := 1
	two := 2

	p1a := &oneA
	p1b := &oneB
	p2 := &two

	if p1a == p1b {
		t.Errorf("shouldn't pointers differ? %v vs %v", p1a, p1b)
	}

	queue.Push(*p1a)
	queue.Push(*p2)

	a := queue.Peek()
	if *a != 1 {
		t.Errorf("peek() expected 1; got %d", *a)
	}

	b := queue.Peek()
	if *b != 1 {
		t.Errorf("peek() expected 1; got %d", *b)
	}

	c := queue.Next()
	if *c != 1 {
		t.Errorf("next() expected 1; got %d", *c)
	}

	d := queue.Next()
	if *d != 2 {
		t.Errorf("next() expected 2; got %d", *d)
	}

	if c != p1a {
		t.Errorf("expected p1a & c pointers to be equal")
	}

	if d != p2 {
		t.Errorf("expected p2 & d pointers to be equal")
	}
}
