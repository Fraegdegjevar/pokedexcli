package pokeapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetLocationAreas(t *testing.T) {
	// Important that we test how we handle output/format from the api.
	// Not the api itself. So we mock-up an api and check we parse responses properly

	// Mock JSON response
	mockJSON := `{
	"next": "https://testpokeapi.localtest/api/v2/location-area/?offset=20&limit=20",
	"previous": null,
	"results": [
		{"name": "test-area-1", "url": "https://testpokeapi.localtest/api/v2/location-area/1/"},
		{"name": "test-area-2", "url": "https://testpokeapi.localtest/api/v2/location-area/2/"}
		]
	}`

	//Test HTTP server - exists in memory. This is a goroutine listening on a random available
	// localhost port. When it starts, go will assign it a url i.e 127.0.0.1:PORT
	// which we capture in server.URL.
	// http.HandlerFunc adapts a function into an http.Handler, which respons to an HTTP request.
	// It writes reply headers and data to a ResponseWriter and returns.
	// A ResponseWriter interface is used by a handler to construct an HTTP response.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Let's log the request coming in
		t.Logf("\n***received request: %s %s***\n", r.Method, r.URL.String())

		//Simulate a few different behaviours
		switch r.URL.Path {
		case "/location-area/":
			// Set headers for realism..
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, mockJSON)
		case "/bad-json/":
			w.Header().Set("Content-Type", "application/json")
			//send malformed response to test GetLocationAreas JSON unmarshal error handling
			fmt.Fprintln(w, `{"results": {}}`)
		default:
			//Error 404 to test GetLocationAreas response error handling
			http.NotFound(w, r)
		}
		//t.Errorf("unexpected path: got %s, want /location-area/", r.URL.Path)

	}))
	defer server.Close()

	conf := &Config{}

	cases := []struct {
		name            string
		path            string
		expectedErr     bool
		expectedResults int
	}{
		{
			name:            "normal response",
			path:            "/location-area/",
			expectedErr:     false,
			expectedResults: 2,
		},
		{
			name:            "malformed json",
			path:            "/bad-json/",
			expectedErr:     true,
			expectedResults: 0,
		},
		{
			name:            "404 not found",
			path:            "/no-exist/",
			expectedErr:     true,
			expectedResults: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			//rebind testcase in our local iteration scope as tests may be run concurrently
			// and overwrite with the last value of tc
			tc := tc
			//ensure url parse properly
			u, err := url.Parse(server.URL + tc.path)
			if err != nil {
				t.Fatalf("error parsing URL %v for test case %v: %v ", tc.path, tc.name, err)
			}
			// Get response
			resp, err := GetLocationAreas(u, conf)
			// First check if our error received matches what we expected
			if (err != nil) != tc.expectedErr {
				t.Errorf("Expected error: %v, got error: %v", tc.expectedErr, err)
			}
			if len(resp.Results) != tc.expectedResults {
				t.Errorf("Expected %v results, got: %v", tc.expectedResults, len(resp.Results))
			}

		})
	}

}
