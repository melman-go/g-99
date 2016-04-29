package g99
import (
	"testing"
	"github.com/bmizerany/assert"
	"github.com/melman-go/g-99/g99"
)

func HandleErr(t *testing.T, err error) {
	if err!=nil {
		t.Error(err.Error())
	}
}

func TestP1(t *testing.T) {
	res1, err := g99.Last([]interface{}{1, 2, 3, 5, 8})
	HandleErr(t, err)
	assert.Equal(t, res1, 8)
	res2, err := g99.Last([]interface{}{1})
	assert.Equal(t, res2, 1)
}
func TestP2(t *testing.T) {
	res1, err := g99.Penultimate([]interface{}{1, 2, 3, 5, 8})
	HandleErr(t, err)
	assert.Equal(t, res1, 5)
	res2, err := g99.Penultimate([]interface{}{1, 2})
	assert.Equal(t, res2, 1)
}
func TestP3(t *testing.T) {
	res1, err := g99.Nth(2, []interface{}{1, 2, 3, 5, 8})
	HandleErr(t, err)
	assert.Equal(t, res1, 2)
	res2, err := g99.Nth(1, []interface{}{1})
	assert.Equal(t, res2, 1)
}
func TestP4(t *testing.T) {
	res1, err := g99.Length([]interface{}{1, 2, 3, 5, 8})
	HandleErr(t, err)
	assert.Equal(t, res1, 5)
}
func TestP5(t *testing.T) {
	res1 := g99.Reverse([]int{1, 2, 3, 5, 8})
	assert.Equal(t, res1, []int{8, 5, 3, 2, 1})
}
func TestP6(t *testing.T) {
	res1, err := g99.IsPalindrome([]interface{}{1, 2, 3, 5, 8})
	HandleErr(t, err)
	assert.Equal(t, res1, false)
	res2, err := g99.IsPalindrome([]interface{}{1, 2, 3, 4, 3, 2, 1})
	assert.Equal(t, res2, true)
}
func TestP7(t *testing.T) {
	res1 := g99.Flatten([]interface{}{1, []interface{}{1, 2, []interface{}{4, 5}}, 3, 5, 8})
	assert.Equal(t, res1, []interface{}{1, 1, 2, 4, 5, 3, 5, 8})
}
func TestP8(t *testing.T) {
	res1 := g99.Compress([]interface{}{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c'})
	assert.Equal(t, res1, []interface{}{'a', 'b', 'c'})
}
func TestP9(t *testing.T) {
	res1 := g99.Pack([]interface{}{'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	assert.Equal(t, res1, []interface{}{[]interface{}{'a', 'a', 'a', 'a'}, []interface{}{'b', 'b'}, []interface{}{'c', 'c', 'c'}})
}
func TestP10(t *testing.T) {
	res1 := g99.Encode([]interface{}{'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	assert.Equal(t, res1, []interface{}{[]interface{}{4, 'a'}, []interface{}{2, 'b'}, []interface{}{3, 'c'}})
}

func TestP11(t *testing.T) {
	res1 := g99.EncodeModified([]interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'c'})
	assert.Equal(t, res1, []interface{}{[]interface{}{4, 'a'}, 'b', []interface{}{3, 'c'}})
}

func TestP12(t *testing.T) {
	res1 := g99.Decode([]interface{}{[]interface{}{4, 'a'}, []interface{}{2, 'b'}, []interface{}{3, 'c'}})
	assert.Equal(t, res1, []interface{}{'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'})
}
func TestP13(t *testing.T) {
	res1 := g99.EncodeDirect([]interface{}{'a', 'a', 'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	assert.Equal(t, res1, []interface{}{[]interface{}{4, 'a'}, []interface{}{2, 'b'}, []interface{}{3, 'c'}})
}

func TestP14(t *testing.T) {
	res1 := g99.Duplicate([]interface{}{'a', 'b', 'c', 'd'})
	assert.Equal(t, res1, []interface{}{'a', 'a', 'b', 'b', 'c', 'c', 'd', 'd'})
}

func TestP15(t *testing.T) {
	res1 := g99.DuplicateN(3, []interface{}{'a', 'b', 'c', 'd'})
	assert.Equal(t, res1, []interface{}{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c', 'd', 'd', 'd'})
}
func TestP16(t *testing.T) {
	res1 := g99.Drop(3, []interface{}{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'})
	assert.Equal(t, res1, []interface{}{'a', 'b', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'})
}
func TestP17(t *testing.T) {
	res1 := g99.Split(3, []interface{}{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'})
	assert.Equal(t, res1, []interface{}{[]interface{}{'a', 'b', 'c'}, []interface{}{'d', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}})
}
func TestP18(t *testing.T) {
	res1 := g99.Slice(3, 7, []interface{}{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'})
	assert.Equal(t, res1, []interface{}{'d', 'e', 'f', 'g'})
}

func TestP19(t *testing.T) {
	res1 := g99.Rotate(3, []interface{}{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'})
	assert.Equal(t, res1, []interface{}{'d', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'a', 'b', 'c'})

	res2 := g99.Rotate(-2, []interface{}{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'})
	assert.Equal(t, res2, []interface{}{'j', 'k', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'})
}

func TestP20(t *testing.T) {
	res1 := g99.RemoveAt(2, []interface{}{'a', 'b', 'c', 'd'})
	assert.Equal(t, res1, []interface{}{[]interface{}{'a', 'c', 'd'}, 'b'})
}

func TestP21(t *testing.T) {
	res1 := g99.InsertAt("new", 1, []interface{}{'a', 'b', 'c', 'd'})
	assert.Equal(t, res1, []interface{}{'a', "new", 'b', 'c', 'd'})
}

func TestP22(t *testing.T) {
	res1 := g99.Range(4, 9)
	assert.Equal(t, res1, []int{4, 5, 6, 7, 8, 9})
}

func TestP23(t *testing.T) {
	res1 := g99.RandomSelect(3, []interface{}{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'})
	assert.Equal(t, len(res1), 3)
}
func TestP24(t *testing.T) {
	res1 := g99.Lotto(6, 49)
	assert.Equal(t, len(res1), 6)
}
func TestP25(t *testing.T) {
	res1 := g99.RandomePermute([]interface{}{'a', 'b', 'c', 'd', 'e', 'f'})
	assert.Equal(t, len(res1), 6)
}
func TestP26(t *testing.T) {
	res1 := g99.Combinations(3, []interface{}{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'})
	assert.Equal(t, len(res1), 220)
}
func TestP27(t *testing.T) {
	res1 := g99.Group3([]int{2, 2, 5}, []interface{}{"Aldo", "Beat", "Carla", "David", "Evi", "Flip", "Gary", "Hugo", "Ida"})
	assert.Equal(t, len(res1), 3)
}
func TestP28(t *testing.T) {
	res1 := g99.LsortFreq([]interface{}{[]interface{}{'a', 'b', 'c'}, []interface{}{'d', 'e'}, []interface{}{'f', 'g', 'h'}, []interface{}{'d', 'e'}, []interface{}{'i', 'j', 'k', 'l'}, []interface{}{'m', 'n'}, []interface{}{'o'}})
	assert.Equal(t, res1, []interface{}{[]interface{}{'i', 'j', 'k', 'l'}, []interface{}{'o'}, []interface{}{'a', 'b', 'c'}, []interface{}{'f', 'g', 'h'}, []interface{}{'d', 'e'}, []interface{}{'d', 'e'}, []interface{}{'m', 'n'}})
}