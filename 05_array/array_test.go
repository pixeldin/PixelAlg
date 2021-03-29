package _5_array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	arr.Insert(uint(6), 999)
	arr.Print()

	arr.InsertToTail(666)
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if nil != err {
			t.Fatal(err)
		}
		arr.Print()
	}
}

func TestFind(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))
}

var SrcArray = []int{1, 2, 3, 3, 4, 4, 5}

func TestSearchByDiv(t *testing.T) {
	div := SearchByDiv(4, SrcArray)
	fmt.Println(div)
	div = SearchByDiv(1, SrcArray)
	fmt.Println(div)
	div = SearchByDiv(5, SrcArray)
	fmt.Println(div)
	div = SearchByDiv(9, SrcArray)
	fmt.Println(div)
}

func TestBinarySearch(t *testing.T) {
	search := BinarySearch(SrcArray, 1, 0, len(SrcArray)-1)
	fmt.Println(search)
	search = BinarySearch(SrcArray, 5, 0, len(SrcArray)-1)
	fmt.Println(search)
}
