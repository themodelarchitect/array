package array

import (
	"testing"
)

func TestArray_New(t *testing.T) {
	arr := NewArray[int]()
	if arr.length != 0 {
		t.FailNow()
	}
}

func BenchmarkArray_New(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewArray[int]()
	}
}

func TestArray_Push(t *testing.T) {
	arr := NewArray[int]()
	for i := 0; i < 10; i++ {
		n := arr.Push(i)
		if n != i+1 {
			t.FailNow()
		}
	}
}

func BenchmarkArray_Push(b *testing.B) {
	arr := NewArray[int]()
	for n := 0; n < b.N; n++ {
		arr.Push(n)
	}
}

func TestArray_Pop(t *testing.T) {
	arr := NewArray[int]()
	for i := 0; i < 10; i++ {
		arr.Push(i)
	}
	n := arr.Pop()
	if n != 9 {
		t.FailNow()
	}
}

func TestArray_Delete(t *testing.T) {
	arr := NewArray[int]()
	for i := 0; i < 10; i++ {
		arr.Push(i)
	}

	arr.Delete(2)
	if len(arr.data) != 9 {
		t.FailNow()
	}
}

func BenchmarkArray_Delete(b *testing.B) {
	arr := NewArray[int]()
	for n := 0; n < b.N; n++ {
		arr.Push(1)
		arr.Delete(0)
	}
}

func TestArray_Reverse(t *testing.T) {
	arr := NewArray[int]()
	for i := 0; i < 10; i++ {
		arr.Push(i)
	}
	arr.Reverse()
	if arr.data[0] != 9 {
		t.FailNow()
	}
}

func BenchmarkArray_Reverse(b *testing.B) {
	arr := NewArray[int]()
	for i := 0; i < 10; i++ {
		arr.Push(i)
	}

	for n := 0; n < b.N; n++ {
		arr.Reverse()
	}
}
