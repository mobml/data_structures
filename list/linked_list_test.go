package list

import (
	"testing"
)

func TestNew(t *testing.T) {
	result := New()

	if result.Len() != 0 {
		t.Errorf("The lenght of List, got: %d, want: %d", result.Len(), 0)
	}
}

func TestPushBack(t *testing.T) {
	list := New()

	list.PushBack(12)

	if list.Len() != 1 {
		t.Errorf("The lenght of list, got: %d, want: %d", list.Len(), 1)
	}

	value, _ := list.Back().Value()

	if value != 12 {
		t.Errorf("The back of list, got: %d, want: %d", value, 12)
	}
}

func TestPushFront(t *testing.T) {
	list := New()

	if list.Front() != nil {
		t.Errorf("The front of list, got: %v, want: %v", list.Front(), nil)
	}

	list.PushBack(1)
	list.PushFront(2)

	if list.Len() != 2 {
		t.Errorf("The lenght of list, got: %d, want: %d", list.Len(), 2)
	}

	value, _ := list.Front().Value()
	if value != 2 {
		t.Errorf("The front of list, got: %d, want: %d", value, 2)
	}

}

func TestNext(t *testing.T) {
	list := New()

	for i := 0; i < 4; i++ {
		list.PushBack(i)
	}

	i := list.Front()

	if i.next.info != 1 {
		t.Errorf("The next of elements, got: %d, want: %d", i.next.info, 1)
	}
}