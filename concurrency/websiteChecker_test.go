package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "invalid url"
}

// test functionality of WebsiteChecker()
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"invalid url",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"invalid url":                false,
	}

	// using dependency injection to test CheckWebsotes() without making real HTTP calls
	got := CheckWebsites(mockWebsiteChecker, websites)

	// checking array equality usign DeepEqual
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// benchmark test to test the speed of WebsiteChecker(), run using `go test -bench=.`
func BenchmarkCheckWebsites(b *testing.B) {
	// create arr of 100 urls
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer() // reset the time of our test before it actually runs
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
