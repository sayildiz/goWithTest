package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// pre 1.22 go would cause a bug when doing loop
	// for _, url would create a single memory entry for url basically update the value of the *url pointer
	// the for loop would finish before the goroutines start so iterating over  {a,b,c} would result to channel receiving {c,c,c}
	// was fixed with go 1.22 2023
	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
