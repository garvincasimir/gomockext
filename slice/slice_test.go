package slice

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
