package concurrency

import "time"

// func
type WebsiteChecker func(string) bool

type result struct {
	// as we don't need either value to be named, each of them is anonymous within the struct
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result) // Create channel usign chan to avoid concurrent writes to results map. 'chan result' - result is the type of channel we create.

	// An operation that does not block in Go will run in a separate process called a goroutine.
	// TO start a new goroutine, turn a function call into a go statement by putting the keyword go in front of it: go doSomething()
	// 1. we often use anonymous functions when we want to start a goroutine
	// 2. all the variables that are available at the point when you declare the anonymous function are also available in the body of the function
	// each iteration of the loop will start a new goroutine, concurrent with the current process (the WebsiteChecker function)
	for _, url := range urls {

		// variable url is reused for each iteration of the for loop - it takes a new value from urls each time.
		// But each of our goroutines have a reference to the url variable - they don't have their own independent copy.
		// So they'll write the value that `url` has at the end of iteration (the last url)
		// To avoid this, pass u to each anonymous function - a copy of the value of url, and so can't be changed.
		// Calling the anonymous function with the url as the argument, we make sure that the value of u is fixed as the value of url
		go func(u string) {
			// results[u] = wc(u) // directly writing to results gives data race
			resultChannel <- result{u, wc(u)} // sending a result struct for each call to wc to the resultChannel with a send to channel statement
		}(url) // anonymous functions executed at the same time that they're declared, using ()
	}

	// `go test -race` to detect race condition

	// solve this data race by coordinating our goroutines using channels. Channels are a
	// Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.

	// goroutines that our for loop started need time to add their result to the results map
	time.Sleep(2 * time.Second)

	// writing to resutls map one at a time (somethign that can't be parallelized),
	// although each of the calls of wc, and each send to the result channel, is happening in parallel inside its own process
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // blocking call, receive expression, assign value from channel to a variable
		results[r.string] = r.bool
	}

	return results
}
