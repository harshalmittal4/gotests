# testing

- Because our ReverseRunes function begins with an upper-case letter, it is exported, and can be used in other packages that import our morestrings package.

- Tests
file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T). The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed (<https://go.dev/doc/code>)
