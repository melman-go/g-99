package util
import (
	"github.com/melman-go/g-99/util"
	"testing"
)

type SortTestEntity struct {
	Id   int
	Name string
}

type TestEntityArray []SortTestEntity

func (this TestEntityArray) Len() int {
	return len(this)
}
func (this TestEntityArray) Swap(a int, b int) {
	this[a], this[b] = this[b], this[a]
}

func (this TestEntityArray) GetKey(pos int) int {
	return this[pos].Id
}

func (this TestEntityArray) Get(pos int) interface{} {
	return this[pos]
}
func (this TestEntityArray) Set(pos int, data interface{}) {
	this[pos] = data.(SortTestEntity)
}

func TestSort(t *testing.T) {
	entity1 := SortTestEntity{Id:68, Name:"a", }
	entity2 := SortTestEntity{Id:57, Name:"b", }
	entity3 := SortTestEntity{Id:23, Name:"c", }
	entity4 := SortTestEntity{Id:38, Name:"d", }
	entity5 := SortTestEntity{Id:73, Name:"e", }
	entity6 := SortTestEntity{Id:36, Name:"f", }
	entity7 := SortTestEntity{Id:43, Name:"g", }
	entity8 := SortTestEntity{Id:12, Name:"h", }
	entity9 := SortTestEntity{Id:62, Name:"i", }
	entity10 := SortTestEntity{Id:99, Name:"j", }
	data := TestEntityArray{entity1, entity2, entity3, entity4, entity5, entity6, entity7, entity8, entity9, entity10}
	util.QuickSort(data)
	t.Log(data)
	data2 := TestEntityArray{entity1, entity2, entity3, entity4, entity5, entity6, entity7, entity8, entity9, entity10}
	util.BubbleSort(data2)
	t.Log(data2)
}