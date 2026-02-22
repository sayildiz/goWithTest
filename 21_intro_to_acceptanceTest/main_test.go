package main

import (
	"testing"
	"time"

	"github.com/quii/go-graceful-shutdown/acceptancetests"
	"github.com/quii/go-graceful-shutdown/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)

	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// check if server works befores TestGracefulShutdown(
	assert.CanGet(t, url)

	// fire off request, before it has chance to respond send SIGTERM
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})

	// without gracful shutdown -> fail
	assert.CanGet(t, url)
	// after interrupt, server should be shutdown, no more request will work
	assert.CantGet(t, url)
}
