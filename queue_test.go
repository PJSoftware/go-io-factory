package factory_test

import (
	"testing"

	factory "github.com/pjsoftware/go-io-factory"
)

func TestQueuePeekEmpty(t *testing.T) {
	const EMPTY = -1
	queue := factory.NewQueue(EMPTY)

	x := queue.Peek()
	if x != EMPTY {
		t.Errorf("expected EMPTY; got %v", x)
	}
}

func TestQueuePeekNext(t *testing.T) {
	const EMPTY = -1
	queue := factory.NewQueue(EMPTY)

	queue.Push(1)
	queue.Push(2)

	a := queue.Peek()
	if a != 1 {
		t.Errorf("peek() expected 1; got %d", a)
	}

	b := queue.Peek()
	if b != 1 {
		t.Errorf("peek() expected 1; got %d", b)
	}

	c := queue.Next()
	if c != 1 {
		t.Errorf("next() expected 1; got %d", c)
	}

	d := queue.Next()
	if d != 2 {
		t.Errorf("next() expected 2; got %d", d)
	}

	e := queue.Next()
	if e != EMPTY {
		t.Errorf("next() expected EMPTY; got %d", e)
	}
}

type testObject struct {
	ident int
	value int
}

var id = 1

func newTestObject(value int) *testObject {
	t := &testObject{ident: id, value: value}
	id++

	return t
}

func TestDefaultAndIsEmpty(t *testing.T) {
	toDefault := newTestObject(-1)
	queue1 := factory.NewQueue(toDefault)
	queue2 := factory.NewQueue[*testObject](nil)

	if !queue1.IsEmpty() {
		t.Errorf("expected queue1 to be empty")
	}
	if !queue2.IsEmpty() {
		t.Errorf("expected queue2 to be empty")
	}

	if queue1.Peek() != toDefault {
		t.Errorf("unexpected default value from queue1")
	}
	if queue2.Peek() != nil {
		t.Errorf("unexpected default value from queue2")
	}

	queue1.Push(newTestObject(2))
	if queue1.IsEmpty() {
		t.Errorf("expected queue to NOT be empty")
	}
}

func TestQueuePeekNextPointers(t *testing.T) {
	queue := factory.NewQueue[*testObject](nil)

	oneA := newTestObject(1)

	oneB := newTestObject(1)

	two := newTestObject(2)

	if *oneA == *oneB {
		t.Errorf("should differ: %+v, %+v", *oneA, *oneB)
	}

	if oneA == oneB {
		t.Errorf("should differ: %+v, %+v", oneA, oneB)
	}

	queue.Push(oneA)
	queue.Push(two)

	a := queue.Peek()
	if a.value != 1 {
		t.Errorf("peek() expected 1; got %d", a.value)
	}

	b := queue.Peek()
	if b.value != 1 {
		t.Errorf("peek() expected 1; got %d", b.value)
	}

	if a != b {
		t.Errorf("shouldn't pointers be equal? %v vs %v", a, b)
	}

	c := queue.Next()
	if c.value != 1 {
		t.Errorf("next() expected 1; got %d", c.value)
	}

	d := queue.Next()
	if d.value != 2 {
		t.Errorf("next() expected 2; got %d", d.value)
	}

	if c != oneA {
		t.Errorf("expected oneA & c pointers to be equal")
	}

	if *d != *two {
		t.Errorf("expected two & d dereferenced pointers to be equal")
	}
}
