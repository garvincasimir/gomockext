package stringext

import "testing"

func TestStartsWithFalseWithNoMatch(t *testing.T) {
	matcher := StartsWith("abc", "Starts with abc")
	if matcher.Matches("def") {
		t.Error("def does not start with abc")
	}
}

func TestStartsWithTrueWithMatch(t *testing.T) {
	matcher := StartsWith("abc", "Starts with abc")
	if !matcher.Matches("abcdef") {
		t.Error("abcdef does start with abc")
	}
}

func TestStartsWithFalseWithNoneString(t *testing.T) {
	matcher := StartsWith("abc", "Starts with abc")
	if matcher.Matches(823) {
		t.Error("Should not match on non string")
	} 
}


func TestEndsWithFalseWithNoMatch(t *testing.T) {
	matcher := EndsWith("abc", "Ends with abc")
	if matcher.Matches("def") {
		t.Error("def does not end with abc")
	}
}

func TestEndsWithTrueWithMatch(t *testing.T) {
	matcher := EndsWith("def", "Ends with abc")
	if !matcher.Matches("abcdef") {
		t.Error("abcdef does end with def")
	}
}

func TestEndsWithFalseWithNoneString(t *testing.T) {
	matcher := EndsWith("abc", "Ends with abc")
	if matcher.Matches(823) {
		t.Error("Should not match on non string")
	} 
}


func TestContainsFalseWithNoMatch(t *testing.T) {
	matcher := Contains("abc", "Contains abc")
	if matcher.Matches("def") {
		t.Error("Does does contain abc")
	}
}

func TestContainsTrueWithMatch(t *testing.T) {
	matcher := Contains("def", "Contains def")
	if !matcher.Matches("abcdef") {
		t.Error("abcdef does contain def")
	}
}

func TestContainsFalseWithNoneString(t *testing.T) {
	matcher := Contains("abc", "Contains abc")
	if matcher.Matches(823) {
		t.Error("Should not match on non string")
	} 
}
