package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPing tests ping handler for the correct response status code, 200 and
// the correct response body, "OK".
func TestPing(t *testing.T) {
	t.Parallel()
	// Initialize a new httptest.ResponseRecorder.
	rr := httptest.NewRecorder()

	// Initialize a new dummy http.Request.
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up basic application struct for testing ping handler
	app := &application{}

	// Call the ping handler function, passing the httptest.ResponseRecorder and http.Request.
	app.ping(rr, r)

	// Call the Result() method on the http.ResponseRecorder to get the http.Response
	// generated by the ping handler.
	rs := rr.Result()

	// We can then examine the http.Response to check that the status code
	// written by the ping handler was 200.
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	// And we can check that the response body written by the ping handler equals "OK".
	defer func() {
		if err := rs.Body.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
