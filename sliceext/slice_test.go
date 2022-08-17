package sliceext

import "testing"

func TestHasItemTrueIfMatch(t *testing.T) {
	matcher := HasItem("abc", "Contains item abc")

	if !matcher.Matches([]string{"abc", "def"}) {
		t.Error("abc is in list")
	}
}

func TestHasItemFalseIfNoMatch(t *testing.T) {
	matcher := HasItem("gh", "Contains item gh")

	if matcher.Matches([]string{"abc", "def"}) {
		t.Error("gh is not in the list")
	}
}

func TestHasItemFalseIfNotListT(t *testing.T) {
	matcher := HasItem("gh", "Contains item gh")

	if matcher.Matches([]int{1, 2}) {
		t.Error("should not match on different type")
	}
}

func TestHasItemsTrueIfMatch(t *testing.T) {
	matcher := HasItems([]string{"abc", "def"}, "Contains item abc and def")

	if !matcher.Matches([]string{"abc", "def"}) {
		t.Error("abc and def are in the list")
	}
}

func TestHasItemsFalseIfNoMatch(t *testing.T) {
	matcher := HasItems([]string{"yyy", "zzzz"}, "Contains item yyy and zzz")

	if matcher.Matches([]string{"abc", "def"}) {
		t.Error("yyy and zzz not in list")
	}
}

func TestHasItemsFalseIfPartialMatch(t *testing.T) {
	matcher := HasItems([]string{"abc", "zzzz"}, "Contains item yyy and zzz")

	if matcher.Matches([]string{"abc", "def"}) {
		t.Error("abc in list but zzz not in list")
	}
}

func TestHasItemsFalseIfNotListT(t *testing.T) {
	matcher := HasItems([]string{"abc", "def"}, "Contains item gh")

	if matcher.Matches([]int{1, 2}) {
		t.Error("should not match on different type")
	}
}

func TestHasItemAtTrueIfAtIndex(t *testing.T) {
	matcher := HasItemAt(9, 3, "there are 4 elements and 9 is the 4th")

	if !matcher.Matches([]int{4, 5, 6, 9}) {
		t.Error("9 should be the 4th element")
	}
}

func TestHasItemAtFalseIfInListButNotAtIndex(t *testing.T) {
	matcher := HasItemAt(9, 2, "there are 4 elements and 9 is the 3rd")

	if matcher.Matches([]int{4, 5, 6, 9}) {
		t.Error("9 is not the 3rd element")
	}
}

func TestHasItemAtFalseIfNotInList(t *testing.T) {
	matcher := HasItemAt(9, 2, "there are 4 elements and 9 is the 3rd")

	if matcher.Matches([]int{4, 5, 6, 7}) {
		t.Error("9 is not in the list")
	}
}

func TestHasItemAtFalseIfNotListT(t *testing.T) {
	matcher := HasItemAt(9, 2, "there are 4 elements and 9 is the 3rd")

	if matcher.Matches([]string{"4", "7", "9", "5"}) {
		t.Error("List not correct type")
	}
}

func TestAnyTrueifAnyMatches(t *testing.T) {
	matcher := Any(func(item int) bool { return item == 9 }, "9 exists in the list")

	if !matcher.Matches([]int{4, 5, 9, 7}) {
		t.Error("9 is in the list")
	}
}

func TestAnyFalseIfListEmpty(t *testing.T) {
	matcher := Any(func(item int) bool { return item == 9 }, "9 exists in the list")

	if matcher.Matches([]int{}) {
		t.Error("List is empty")
	}
}

func TestAnyFalseIfNoMatches(t *testing.T) {
	matcher := Any(func(item int) bool { return item == 9 }, "9 exists in the list")

	if matcher.Matches([]int{4, 5, 6, 7}) {
		t.Error("9 is not in the list")
	}
}

func TestAnyFalseIfNoListT(t *testing.T) {
	matcher := Any(func(item int) bool { return item == 9 }, "9 exists in the list")

	if matcher.Matches([]string{"2", "9", "4"}) {
		t.Error("not a list of int")
	}
}

func TestAllTrueIfAllMatch(t *testing.T) {
	matcher := All(func(item int) bool { return item > 3 }, "all items more than 3")

	if !matcher.Matches([]int{4, 5, 9, 7}) {
		t.Error("all items in the list are more than 3")
	}
}

func TestAllFalseIfNotAllMatch(t *testing.T) {
	matcher := All(func(item int) bool { return item > 3 }, "all items more than 3")

	if matcher.Matches([]int{4, 5, 9, 7, 3}) {
		t.Error("all items in the list are not more than 3")
	}
}

//https://github.com/microsoft/referencesource/blob/master/System.Core/System/Linq/Enumerable.cs#L1305
func TestAllTrueIfListEmpty(t *testing.T) {
	matcher := All(func(item int) bool { return item > 3 }, "all items more than 3")

	if !matcher.Matches([]int{}) {
		t.Error("Empty list still matches")
	}
}

func TestAllFalseIfNotListT(t *testing.T) {
	matcher := All(func(item int) bool { return item > 3 }, "all items more than 3")

	if matcher.Matches([]string{"4", "5"}) {
		t.Error("This is not a list of int")
	}
}
