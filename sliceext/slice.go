package sliceext

import mockext "github.com/garvincasimir/gomockext"

// HasItem matcher returns true if actual is a []T and it contains item
//
// Example usage:
//   HasItem("abc", "Contains item abc")
func HasItem[T comparable](item T, message string) mockext.CustomMatcher {
	return HasItems([]T{item}, message)
}

// HasItemAt matcher returns true actual is a []T and it contains item at a specific index
// If the index is out of bounds it will return false
//
// Example usage:
//   HasItemAt("abc", "Contains item abc at index 1", 1)
func HasItemAt[T comparable](item T, index int, message string) mockext.CustomMatcher {
	return mockext.Match(func(got interface{}) bool {
		if s, ok := got.([]T); ok && len(s) > index {
			return s[index] == item
		}
		return false
	}, message)
}

// HasItems matcher returns true if actual is a []T subset is contained within []T
//
// Example usage:
//   HasItems([]string{"abc", "def"}, "Contains item abc, def")
func HasItems[T comparable](subset []T, message string) mockext.CustomMatcher {
	return mockext.Match(func(got interface{}) bool {
		if s, ok := got.([]T); ok {
			set := map[T]bool{}
			for _, v := range s {
				set[v] = true
			}

			for _, v := range subset {
				if !set[v] {
					return false
				}
			}
			return true
		}
		return false
	}, message)
}

// Any matcher returns true if actual is a []T and the checkFuc returns true for any of the items
//
// Example usage:
//   Any(func(u user)bool { return u.name == "joe" }, "One of the users is Joe")
func Any[T any](checkFunc func(T) bool, message string) mockext.CustomMatcher {
	return mockext.Match(func(got interface{}) bool {
		if s, ok := got.([]T); ok {
			for _, v := range s {
				if checkFunc(v) {
					return true
				}
			}
		}
		return false
	}, message)
}

// All matcher returns true if actual is a []T and the checkFuc returns true for all of the items
//
// Example usage:
//   All(func(u user)bool { return u.age > 20" }, "All users over 20")
func All[T any](checkFunc func(T) bool, message string) mockext.CustomMatcher {
	return mockext.Match(func(got interface{}) bool {
		if s, ok := got.([]T); ok {
			for _, v := range s {
				if !checkFunc(v) {
					return false
				}
			}
		} else {
			return false
		}
		return true
	}, message)
}
