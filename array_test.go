package array

import (
	"log"
	"testing"
)

func TestArray_1(t *testing.T) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)

	for i := 0; i < 3; i++ {
		val := arr.Lookup(i)
		if val != i+1 {
			t.FailNow()
		}
	}
}

func TestArray_2(t *testing.T) {
	expected := []string{"a", "b", "c"}

	arr := New[string]()
	arr.Push("a")
	arr.Push("b")
	arr.Push("c")

	for i := 0; i < 3; i++ {
		val := arr.Lookup(i)
		if val != expected[i] {
			t.FailNow()
		}
	}
}

func TestArray_Merge(t *testing.T) {
	arr1 := New[int]()
	arr2 := New[int]()

	arr1.Push(1)
	arr1.Push(2)
	arr1.Push(3)

	arr2.Push(4)
	arr2.Push(5)
	arr2.Push(6)

	mergedArray := arr1.Merge(arr2)

	for i := 0; i < 6; i++ {
		n := mergedArray.Lookup(i)
		t.Log(n)
		if n != i+1 {
			t.FailNow()
		}
	}
}

func TestArray_3(t *testing.T) {
	expected := []string{"c", "b", "a"}

	arr := New[string]()
	arr.Push("a")
	arr.Push("b")
	arr.Push("c")
	arr.Reverse()

	for i := 0; i < 3; i++ {
		val := arr.Lookup(i)
		if val != expected[i] {
			t.FailNow()
		}
	}
}

func TestArray_4(t *testing.T) {
	arr := New[int]()

	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	value := arr.Pop()
	if value != 3 {
		t.FailNow()
	}
	value = arr.Pop()
	if value != 2 {
		t.FailNow()
	}
	value = arr.Pop()
	if value != 1 {
		t.FailNow()
	}
	count := arr.Length()
	if count != 0 {
		t.FailNow()
	}
}

func TestArray_Copy(t *testing.T) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arrCopy := arr.Copy()

	for i := 0; i < 3; i++ {
		x := arr.Lookup(i)
		y := arrCopy.Lookup(i)
		log.Println(x, y)
		if x != y {
			t.FailNow()
		}
	}
}

func TestArray_Copy2(t *testing.T) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arrCopy := arr.Copy()

	for i := 0; i < arr.length; i++ {
		arrCopy.Set(i, i+3)
	}

	for i := 0; i < arr.length; i++ {
		x := arr.Lookup(i)
		y := arrCopy.Lookup(i)
		log.Println(x, y)
	}
}

func TestArray_Set(t *testing.T) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)

	arr.Set(0, 4)
	arr.Set(1, 5)
	arr.Set(2, 6)

	for i := 0; i < 3; i++ {
		x := arr.Lookup(i)
		log.Println(x)
		if x != (i + 4) {
			t.FailNow()
		}
	}

}

func TestArray_Clear(t *testing.T) {
	arr := New[int]()
	for i := 0; i < 10; i++ {
		arr.Push(i)
	}

	if arr.Length() != 10 {
		t.FailNow()
	}

	arr.Clear()

	if arr.Length() > 0 {
		t.FailNow()
	}
}

func BenchmarkArray_Push(b *testing.B) {
	arr := New[int]()
	for n := 0; n < b.N; n++ {
		arr.Push(n)
	}
}

func BenchmarkArray_Copy(b *testing.B) {
	arr := New[int]()
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	for n := 0; n < b.N; n++ {
		arr.Copy()
	}
}
