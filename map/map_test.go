package mapext

import "testing"

func TestHasKeyTrueIfMatch(t *testing.T) {
	matcher := HasKey[string, int]("abc", "Contains key abc")

	if !matcher.Matches(map[string]int{"abc": 1, "def": 2}) {
		t.Error("Key abc exists in map")
	}
}

func TestHasKeyFalseIfNoMatch(t *testing.T) {
	matcher := HasKey[string, int]("abc", "Contains key abc")

	if matcher.Matches(map[string]int{"xyz": 1, "def": 2}) {
		t.Error("Key abc does not exist in map")
	}
}

func TestHasKeyFalseIfNoMapKV(t *testing.T) {
	matcher := HasKey[string, int]("abc", "Contains key abc")

	if matcher.Matches(map[string]string{"abc": "1", "def": "2"}) {
		t.Error("Wrong map type")
	}
}

func TestHasKeysTrueIfMatch(t *testing.T) {
	matcher := HasKeys[string, int]([]string{"abc", "def"}, "Contains key abc and def")

	if !matcher.Matches(map[string]int{"abc": 1, "def": 2, "xyz": 3}) {
		t.Error("Keys exist in map")
	}
}

func TestHasKeysFalseIfPartialMatch(t *testing.T) {
	matcher := HasKeys[string, string]([]string{"abc", "def"}, "Contains keys abc and def")

	if matcher.Matches(map[string]int{"xyz": 1, "def": 2}) {
		t.Error("Key def exists but abc doesn't")
	}
}

func TestHasKeysFalseIfNoMatch(t *testing.T) {
	matcher := HasKeys[string, string]([]string{"abc", "def"}, "Contains keys abc and def")

	if matcher.Matches(map[string]int{"xyz": 1, "lmn": 2}) {
		t.Error("Keys def an abc dont exist in map")
	}
}

func TestHasKeysFalseIfNoMapKV(t *testing.T) {
	matcher := HasKeys[string, int]([]string{"abc", "def"}, "Contains keys abc and def")

	if matcher.Matches(map[string]string{"abc": "1", "def": "2"}) {
		t.Error("Wrong map type")
	}
}

func TestHasItemTrueIfMatch(t *testing.T) {
	matcher := HasItem("def", 2, "Contains key abc")

	if !matcher.Matches(map[string]int{"abc": 1, "def": 2}) {
		t.Error("Key abc exists in map")
	}
}

func TestHasItemFalseIfNoMatch(t *testing.T) {
	matcher := HasItem("def", 2, "Contains key abc")

	if matcher.Matches(map[string]int{"xyz": 1, "abc": 2}) {
		t.Error("Item def:2 does not exist in map")
	}
}

func TestHasItemFalseIfNoMapKV(t *testing.T) {
	matcher := HasItem("def", 2, "Contains key abc")

	if matcher.Matches(map[string]string{"abc": "1", "def": "2"}) {
		t.Error("Wrong map type")
	}
}

func TestHasItemsTrueIfMatch(t *testing.T) {
	matcher := HasItems(map[string]int{"abc": 1, "def": 2}, "Contains items abc:1 and def:2")

	if !matcher.Matches(map[string]int{"abc": 1, "def": 2, "xyz": 3}) {
		t.Error("Items abc:1 and def:2 exist in map")
	}
}

func TestHasItemsFalseIfPartialMatch(t *testing.T) {
	matcher := HasItems(map[string]int{"abc": 1, "def": 2}, "Contains items abc:1 and def:2")

	if matcher.Matches(map[string]int{"xyz": 1, "def": 2}) {
		t.Error("Item def:2 exists but abc:1 doesn't")
	}
}

func TestHasItemsFalseIfNoMatch(t *testing.T) {
	matcher := HasItems(map[string]int{"abc": 1, "def": 2}, "Contains items abc:1 and def:2")

	if matcher.Matches(map[string]int{"xyz": 1, "lmn": 2}) {
		t.Error("Keys def an abc dont exist in map")
	}
}

func TestHasItemsFalseIfNoMapKV(t *testing.T) {
	matcher := HasItems(map[string]int{"abc": 1, "def": 2}, "Contains items abc:1 and def:2")

	if matcher.Matches(map[string]string{"abc": "1", "def": "2"}) {
		t.Error("Wrong map type")
	}
}

func TestAnyTrueIfAnyMatches(t *testing.T) {
	matcher := Any(func(k string, v int) bool { return v == 9 }, "Any values equal 9")

	if !matcher.Matches(map[string]int{"abc": 1, "def": 9, "xyz": 3}) {
		t.Error("9 Exists as a value in the map")
	}
}

func TestAnyFalseIfNoMatches(t *testing.T) {
	matcher := Any(func(k string, v int) bool { return v > 10 }, "Any values > 10")

	if matcher.Matches(map[string]int{"abc": 1, "def": 9, "xyz": 3}) {
		t.Error("None of the values are more than 10")
	}
}

func TestAnyFalseIfListEmpty(t *testing.T) {
	matcher := Any(func(k string, v int) bool { return v > 10 }, "Any values > 10")

	if matcher.Matches(map[string]int{}) {
		t.Error("List is empty")
	}
}

func TestAnyFalseIfNoMapKV(t *testing.T) {
	matcher := Any(func(k string, v int) bool { return v > 10 }, "Any values > 10")

	if matcher.Matches(map[string]string{"abc": "11", "def": "9"}) {
		t.Error("Wrong map type")
	}
}

func TestAllTrueIfAllMatch(t *testing.T) {
	matcher := All(func(k string, v int) bool { return v > 9 }, "All values > 9")

	if !matcher.Matches(map[string]int{"abc": 11, "def": 91, "xyz": 13}) {
		t.Error("All values in the map are > 9")
	}
}

func TestAllFalseIfNoAllMatch(t *testing.T) {
	matcher := All(func(k string, v int) bool { return v > 10 }, "All values > 10")

	if matcher.Matches(map[string]int{"abc": 1, "def": 22, "xyz": 93}) {
		t.Error("Not all the values match")
	}
}

func TestAllTrueIfListEmpty(t *testing.T) {
	matcher := All(func(k string, v int) bool { return v > 10 }, "Any values > 10")

	if !matcher.Matches(map[string]int{}) {
		t.Error("List is empty")
	}
}

func TestAllFalseIfNoMapKV(t *testing.T) {
	matcher := All(func(k string, v int) bool { return v > 9 }, "All values > 9")

	if matcher.Matches(map[string]string{"abc": "11", "def": "91", "xyz": "13"}) {
		t.Error("Wrong map type")
	}
}
