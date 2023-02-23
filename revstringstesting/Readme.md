# testing

- Because our ReverseRunes function begins with an upper-case letter, it is exported, and can be used in other packages that import our morestrings package.

- Tests
go test -v

file with a name ending in _test.go (tells the go tools that this is a testing file and not to be used for build that are not tests, and whenever running tests use this source file) that contains functions named TestXXX (tells go tools this function need to be included as a test case).  *testing.T - test tool expects every test function to accept this, this is what it passes in so that we can tell if the test is passing, whether we want to skip the test. TestXXX (instead of defining the test in main()) runs, it's a function that is being called by some other program.

When calling go test, Go tool take all _test.go source files, uses them to build a new binary that'll run, and then store the binary in a temporary directory and run. It's coming with its own main() function (equivalent of it). At the end the Go tool will clean up the binary or any old artifacts it has.

Need to have multiple tests using a similar setup, or sharing some common funtionality throughout the test, easier to write multiple test cases within the same funtion and sharing some common code - table driven tests (design pattern) - looks like a array of test cases where each test case is a struct with the input and expected output, on which we iterate using a for loop and run the test for each one.

We can use t.Run() to run a subtest for each test case. When using subtests we can use t.Fatal() but without it we need to use t.Errorf(), because if we use t.Fatal() the rest of the test cases in the table after one of them fails will stop.

Whenever signalling a failure in the test case, instead of panic (that will stop the execution of what we are running), we signal that there is failure - t.Errorf(), t.Fail() etc - notify the test tool/ test framework that this test failed.
