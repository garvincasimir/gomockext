package example

import (
	"strings"
	"testing"

	"github.com/garvincasimir/gomockext/mapext"
	"github.com/garvincasimir/gomockext/sliceext"
	"github.com/garvincasimir/gomockext/stringext"
	gomock "github.com/golang/mock/gomock"
)

func TestArg1StringStartsWith(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithString(stringext.StartsWith("Foo", "Arg should start with Foo")).Times(1)
	fooBar.WithString("FooBar")
}

func TestArg1StringEndsWith(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithString(stringext.EndsWith("Bar", "Arg should end with Bar")).Times(1)
	fooBar.WithString("FooBar")
}

func TestArg1StringContains(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithString(stringext.Contains("ooBar", "Arg should contain ooBar")).Times(1)
	fooBar.WithString("FooBar")
}

func TestArg1SliceHasItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithSlice(sliceext.HasItem("foo", "Arg should contain item foo")).Times(1)
	fooBar.WithSlice([]string{"foo", "bar", "abc"})
}

func TestArg1SliceHasItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithSlice(sliceext.HasItems([]string{"foo", "bar"}, "Arg should contain items foo and bar")).Times(1)
	fooBar.WithSlice([]string{"foo", "bar", "abc"})
}

func TestArg1SliceHasItemAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithSlice(sliceext.HasItemAt("bar", 1, "bar is at position 2")).Times(1)
	fooBar.WithSlice([]string{"foo", "bar", "abc"})
}

func TestArg1SliceAny(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithSlice(sliceext.Any(func(item string) bool { return item == "bar" }, "bar exists in the list")).Times(1)
	fooBar.WithSlice([]string{"foo", "bar", "abc"})
}

func TestArg1SliceAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithSlice(sliceext.All(func(item string) bool { return strings.HasPrefix(item, "f") }, "All start with f")).Times(1)
	fooBar.WithSlice([]string{"foo", "fun", "fooBar"})
}

func TestArg1MapHasKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	//Can't infer the type of value here but the library needs to know
	fooBar.EXPECT().WithMap(mapext.HasKey[string, int]("fun", "Map contains the key foo")).Times(1)
	fooBar.WithMap(map[string]int{"foo": 1, "fun": 2, "fooBar": 3})
}

func TestArg1MapHasKeys(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	//Can't infer the type of value here but the library needs to know
	fooBar.EXPECT().WithMap(mapext.HasKeys[string, int]([]string{"fun", "foo"}, "Map contains the keys foo and fun")).Times(1)
	fooBar.WithMap(map[string]int{"foo": 1, "fun": 2, "fooBar": 3})
}

func TestArg1MapHasItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithMap(mapext.HasItem("fun", 2, "Map contains the item fun:2")).Times(1)
	fooBar.WithMap(map[string]int{"foo": 1, "fun": 2, "fooBar": 3})
}

func TestArg1MapHasItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithMap(mapext.HasItem("fun", 2, "Map contains the item fun:2")).Times(1)
	fooBar.WithMap(map[string]int{"foo": 1, "fun": 2, "fooBar": 3})
}

func TestArg1MapAny(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithMap(mapext.Any(func(i string, v int) bool { return v > 10 }, "contains values that are > 10")).Times(1)
	fooBar.WithMap(map[string]int{"foo": 1, "fun": 22, "fooBar": 3})
}

func TestArg1MapAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithMap(mapext.All(func(i string, v int) bool { return v < 23 }, "all values are < 23")).Times(1)
	fooBar.WithMap(map[string]int{"foo": 1, "fun": 22, "fooBar": 3})
}
