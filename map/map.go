package mapext

import (
	mockext "github.com/garvincasimir/gomockext"
)

// HasKey matcher returns true if actual is a map[K]V and all the keys exist in map[K]V
//
// Example usage:
//   HasKey("id1", "Has joe")
func HasKey[K comparable, V any](key K, message string) mockext.CustomMatcher {
	return HasKeys[K,V]([]K{key}, message)
}

// HasKeys matcher returns true if actual is a map[K]V and all the keys exist in map[K]V
//
// Example usage:
//   HasKeys([]string{"id1", "id2"}], "Has keys id1 and id2")
func HasKeys[K comparable, V any](keys []K, message string) mockext.CustomMatcher {
	return mockext.Match(func(got any) bool {
		if m, ok := got.(map[K]V); ok {
			for _, k := range keys {
				if _, ok := m[k]; !ok {
					return false
				}
			}
		} else {
			return false
		}

		return true
	}, message)
}

// HasItem matcher returns true if actual is a map[K]V and the key and value matches any item in map[K]V
//
// Example usage:
//   HasItem("id1","joe", "Has joe")
func HasItem[K comparable, V comparable](key K, val V, message string) mockext.CustomMatcher {
	return HasItems(map[K]V{key:val}, message)
}


// HasItems matcher returns true if actual is a map[K]V and the items is a subset of map[K]V
//
// Example usage:
//   HasItems(map[string]string{"id1":"joe", "id2":"lisa"}, "Has joe and lisa")
func HasItems[K comparable, V comparable](items map[K]V, message string) mockext.CustomMatcher {
	return mockext.Match(func(got any) bool {
		if gotm, ok := got.(map[K]V); ok {
			for ik, iv := range items {
				if v, ok := gotm[ik]; !(ok && v == iv) {
					return false
				}
			}
		} else {
			return false
		}
		return true
	}, message)
}


// All matcher returns true if actual is a map[K]V and the checkFunc returns true for any of the items
//
// Example usage:
//   All(func(uid string, u User)bool { return u.name == "joe" }, "Any user named joe")
func Any[K comparable, V any](checkFunc func(K, V)bool , message string) mockext.CustomMatcher {
	return mockext.Match(func(got any) bool {
		if m, ok := got.(map[K]V); ok {
			for k, v := range m {
				if checkFunc(k, v) {
					return true
				}
			}
		}
		return false
	}, message)
}

// All matcher returns true if actual is a map[K]V and the checkFunc returns true for all of the items
//
// Example usage:
//   All(func(uid string, u User)bool { return u.age > 20" }, "All users over 20")
func All[K comparable, V any](checkFunc func(K, V)bool , message string) mockext.CustomMatcher {
	return mockext.Match(func(got any) bool {
		if m, ok := got.(map[K]V); ok {
			for k, v := range m {
				if checkFunc(k, v) {
					return true
				}
			}
		}
		return false
	}, message)
}