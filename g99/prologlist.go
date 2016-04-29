package g99
import (
	"sort"
	"math/rand"
)


//1.01 (*) Find the last element of a list.
func Last(list []interface{}) (interface{}, error) {
	len, err := GetListLen(list)
	if err!=nil {
		return nil, err
	}
	switch len{
	case 1:return list[0], nil
	default :return list[len-1], nil
	}
}

//1.02 (*) Find the last but one element of a list.
func Penultimate(list []interface{}) (interface{}, error) {
	len, err := GetListLen(list)
	if err!=nil {
		return nil, err
	}
	if len==1 {
		return nil, G99Error{Msg:"list has only one element"}
	}
	return list[len-2], nil
}
//1.03 (*) Find the Kth element of a list.
func Nth(n int, list []interface{}) (interface{}, error) {
	len, err := GetListLen(list)
	if err!=nil {
		return 0, err
	}
	if len<n {
		return 0, G99Error{Msg:"list has only one element"}
	}
	return list[n-1], nil
}
//1.04 (*) Find the number of elements of a list.
func Length(list []interface{}) (int, error) {
	len, err := GetListLen(list)
	if err!=nil {
		return 0, err
	}
	return len, nil
}
//1.05 (*) Reverse a list.
func Reverse(list []int) ([]int) {
	param := list[:]
	sort.Sort(sort.Reverse(sort.IntSlice(param)))
	return param
}
//1.06 (*) Find out whether a list is a palindrome.
func IsPalindrome(list []interface{}) (bool, error) {
	len, err := GetListLen(list)
	if err!=nil {
		return false, err
	}
	for i := 0; i < len/2; i++ {
		if list[i]!=list[len-1-i] {
			return false, err
		}
	}
	return true, err
}
//1.07 (**) Flatten a nested list structure.
func Flatten(list []interface{}) []interface{} {
	res := []interface{}{}
	res =flattenAppend(list, res)
	return res
}

func flattenAppend(list []interface{}, res []interface{}) []interface{} {
	for _, obj := range list {
		switch obj.(type){
			case []interface{} : res = flattenAppend(obj.([]interface{}), res)
			default: res = append(res, obj)
		}
	}
	return res
}
//1.08 (**) Eliminate consecutive duplicates of list elements.
func Compress(list []interface{}) []interface{} {
	res := []interface{}{}
	temp := map[interface{}]bool{}
	for _, obj := range list {
		if temp[obj]!=true {
			res = append(res, obj)
			temp[obj]=true
		}
	}
	return res
}
//1.09 (**) Pack consecutive duplicates of list elements into sublists.
func Pack(list []interface{}) []interface{} {
	res := []interface{}{}
	seq := []interface{}{}
	temp := map[interface{}]int{}
	for _, obj := range list {
		if temp[obj]==0 {
			seq = append(seq, obj)
		}
		temp[obj] = temp[obj]+1
	}
	for _, target := range seq {
		group := []interface{}{}
		for i := 0; i < temp[target]; i++ {
			group = append(group, target)
		}
		res =append(res, group)
	}
	return res
}
//1.10 (*) Run-length encoding of a list.
func Encode(list []interface{}) []interface{} {
	temp := Pack(list)
	res := []interface{}{}
	for _, group := range temp {
		res = append(res, []interface{}{len(group.([]interface{})), group.([]interface{})[0]})
	}
	return res
}
//1.11 (*) Modified run-length encoding.
func EncodeModified(list []interface{}) []interface{} {
	temp := Encode(list)
	res := []interface{}{}
	for _, group := range temp {
		if group.([]interface{})[0]==1 {
			res = append(res, group.([]interface{})[1])
		}else {
			res = append(res, group)
		}
	}
	return res
}
//1.12 (**) Decode a run-length encoded list.
func Decode(list []interface{}) []interface{} {
	res := []interface{}{}
	for _, group := range list {
		for i := 0; i<group.([]interface{})[0].(int); i++ {
			res = append(res, group.([]interface{})[1])
		}
	}
	return res
}
//1.13 (**) Run-length encoding of a list (direct solution).
func EncodeDirect(list []interface{}) []interface{} {
	res := []interface{}{}
	temp := map[interface{}]int{}
	seq := []interface{}{}
	for _, obj := range list {
		if temp[obj]==0 {
			seq = append(seq, obj)
		}
		temp[obj] = temp[obj]+1
	}
	for _, target := range seq {
		group := []interface{}{temp[target], target}
		res =append(res, group)
	}
	return res
}

//1.14 (*) Duplicate the elements of a list.
func Duplicate(list []interface{}) []interface{} {
	res := []interface{}{}
	for _, obj := range list {
		res = append(res, obj)
		res = append(res, obj)
	}
	return res
}

//1.15 (**) Duplicate the elements of a list a given number of times.
func DuplicateN(count int, list []interface{}) []interface{} {
	res := []interface{}{}
	for _, obj := range list {
		for i := 0; i < count; i++ {
			res = append(res, obj)
		}
	}
	return res
}

//1.16 (**) Drop every Nth element from a list.
func Drop(target int, list []interface{}) []interface{} {
	return append(list[:target-1],list[target:]...)
}
//1.17 (*) Split a list into two parts.
func Split(target int, list []interface{}) []interface{} {
	return []interface{}{list[0:target], list[target: len(list)]}
}

//1.18 (**) Extract a slice from a list.
func Slice(start int, end int, list []interface{}) []interface{} {
	return list[start:end]
}

//1.19 (**) Rotate a list N places to the left.
func Rotate(index int, list []interface{}) []interface{} {
	if index ==0 {
		return list
	}
	len, _ := GetListLen(list)
	res := []interface{}{}
	for i := 0; i<len; i++ {
		if index>0 {
			if i+index<len {
				res = append(res, list[i+index])
			}else {
				res = append(res, list[i+index-len])
			}
		}else {
			if i+index>=0 {
				res = append(res, list[i+index])
			}else {
				res = append(res, list[len+index+i])
			}
		}
	}
	return res
}
//1.20 (*) Remove the Kth element from a list.
func RemoveAt(target int, list []interface{}) []interface{} {
	removed := list[target-1]
	return []interface{}{Drop(target, list), removed}
}
//1.21 (*) Insert an element at a given position into a list.
func InsertAt(item interface{}, index int, list []interface{}) []interface{} {
	res := []interface{}{}
	for i := 0; i<len(list)+1; i++ {
		if i<index {
			res = append(res, list[i])
		}else if i==index {
			res = append(res, item)
		}else {
			res = append(res, list[i-1])
		}
	}
	return res
}
//1.22 (*) Create a list containing all integers within a given range.
func Range(start int, end int) []int {
	res := []int{}
	for i := start; i<=end; i++ {
		res = append(res, i)
	}
	return res
}
//1.23 (**) Extract a given number of randomly selected elements from a list.
func RandomSelect(count int, list []interface{}) []interface{} {
	res := []interface{}{}
	for i := 0; i<count; i++ {
		var target interface{}
		index := rand.Intn(len(list))
		temp := RemoveAt(index, list)
		list, target = temp[0].([]interface{}), temp[1]
		res = append(res, target)
	}
	return res
}
//1.24 (*) Lotto: Draw N different random numbers from the set 1..M.
func Lotto(count int, max int) []int {
	res := []int{}
	for i := 0; i<count; i++ {
		res = append(res, rand.Intn(max))
	}
	return res
}
//1.25 (*) Generate a random permutation of the elements of a list.
func RandomePermute(list []interface{}) []interface{} {
	for i:=0;i<len(list);i++{
		index := rand.Intn(len(list))
		list[0],list[index] = list[index],list[0]
	}
	return list
}
//1.26 (**) Generate the combinations of K distinct objects chosen from the N elements of a list.
func Combinations(count int, list []interface{}) []interface{} {
//	res := []interface{}

	return nil
}

func combinationsSub(count int,list []interface{},res []interface{})[]interface{}{
	for i:=0;i<len(list)-count;i++{

	}
	if len(list)-count>1{
//		list = combinationsSub(count,list[1:len(list)])
	}
	return nil
}
//1.27 (**) Group the elements of a set into disjoint subsets.
func Group3(params []int,list []interface{}) []interface{} {
	return nil
}
//1.28 (**) Sorting a list of lists according to length of sublists.
func Lsout(list []interface{}) []interface{}{
	res := []interface{}{}
//	for _,group :=range list{
//
//	}
//	Test()
	return res
}
func LsortFreq(list []interface{}) []interface{} {
	return nil
}