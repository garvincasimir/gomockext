package gomockext

type MatchFunc = func(interface{}) bool

type CustomMatcher struct {
	IsMatch MatchFunc
	Message string
}

func (c CustomMatcher) Matches(got interface{}) bool {
	return c.IsMatch(got)
}

func (c CustomMatcher) String() string {
	return c.Message
}

func Match(isMatch MatchFunc, message string) CustomMatcher {
	return CustomMatcher{
		IsMatch: isMatch,
		Message: message,
	}
}
