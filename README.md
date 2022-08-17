# gomockext
gomockext is a library which adds extensions useful for simplifying testing using [gomock](https://github.com/golang/mock)

## Installation


To get the latest released version use:

### Go 1.18+

```bash
go get github.com/garvincasimir/gomockext@v1.0.0
```
## Usage

### Custom Matcher
You can create a custom matcher to meet more specific conditions by calling the `gomockext.Match` function. 

```go
// StartsWith returns a matcher which checks if a string starts with a specific prefix
func StartsWith(prefix, message string) mockext.CustomMatcher {
	return mockext.Match(func(got any) bool {
		if str, ok := got.(string); ok {
			return strings.HasPrefix(str, prefix)
		}
		return false
	}, message)
}

func TestArg1StringStartsWith(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithString(StartsWith("Foo", "Arg should start with Foo")).Times(1)
	fooBar.WithString("FooBar")
}
```

### Built-In Matcher
The library comes with several built in matchers. Please take a look at the [examples](https://github.com/garvincasimir/gomockext/tree/main/example) folder for usage examples. 

```go
// StartsWith returns a matcher which checks if a string starts with a specific prefix
func TestArg1StringStartsWith(t *testing.T) {
	ctrl := gomock.NewController(t)
	fooBar := NewMockFooBar(ctrl)

	fooBar.EXPECT().WithString(stringext.StartsWith("Foo", "Arg should start with Foo")).Times(1)
	fooBar.WithString("FooBar")
}
```

## Contributing
Please feel free to submit bugs, pull requests and suggestions for improving the library. The repo is configured to create a custom [Codespace](https://github.com/features/codespaces) with everything you need to start contributing