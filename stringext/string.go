package stringext

import (
	mockext "github.com/garvincasimir/gomockext"
	"strings"
)

type StrCheckFunc = func(string, string) bool

// StringsMatcher returns a matcher that calls the supplied strings.[funcName]
//  actual string given the input as the second parameter. Returns false if actual is not a string
func StringsMatcher(input, message string, checkFunc StrCheckFunc) mockext.CustomMatcher {
	return mockext.Match(func(got any) bool {
		if str, ok := got.(string); ok {
			return checkFunc(str, input)
		}
		return false
	}, message)
}

// StartsWith matcher returns true if actual is a string and it starts with the prefix
//
// Example usage:
//   StartsWith("abc", "Starts with value abc")
func StartsWith(prefix, message string) mockext.CustomMatcher {
	return StringsMatcher(prefix, message, strings.HasPrefix)
}

// EndsWith matcher returns true if actual is a string and it ends with the suffix
//
// Example usage:
//   EndsWith("abc", "Ends with value abc")
func EndsWith(suffix, message string) mockext.CustomMatcher {
	return StringsMatcher(suffix, message, strings.HasSuffix)
}

// Contains matcher returns true if the actual string contains the substr
//
// Example usage:
//   Contains("abc", "Contains value abc")
func Contains(substr, message string) mockext.CustomMatcher {
	return StringsMatcher(substr, message, strings.Contains)
}
