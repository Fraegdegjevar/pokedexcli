package pokeapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRequestLocationAreas(t *testing.T) {
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

	cases := []struct {
		name            string
		URL             string
		path            string
		expectedErr     bool
		expectedResults int
	}{
		{
			name:            "normal response",
			URL:             server.URL,
			path:            "/location-area/",
			expectedErr:     false,
			expectedResults: 2,
		},
		{
			name:            "malformed json",
			URL:             server.URL,
			path:            "/bad-json/",
			expectedErr:     true,
			expectedResults: 0,
		},
		{
			name:            "404 not found",
			URL:             server.URL,
			path:            "/no-exist/",
			expectedErr:     true,
			expectedResults: 0,
		},
		{
			name: "no response from api",
			// Note this url is a documentation/test only IP that is guaranteed to be non-routable
			// according to RFC 5737 and will always fail to connect.
			URL:             "http://192.0.2.1:12345",
			path:            "/doesnt-matter/",
			expectedErr:     true,
			expectedResults: 0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			//rebind testcase in our local iteration scope as tests may be run concurrently
			// and overwrite with the last value of tc
			tt := tt
			//ensure url parse properly
			u, err := url.Parse(tt.URL + tt.path)
			if err != nil {
				t.Fatalf("error parsing URL %v for test case %v: %v ", tt.path, tt.name, err)
			}
			// Get response
			resp, err := RequestLocationAreas(u)
			// First check if our error received matches what we expected
			if (err != nil) != tt.expectedErr {
				t.Errorf("Expected error: %v, got error: %v", tt.expectedErr, err)
			}
			if len(resp.Results) != tt.expectedResults {
				t.Errorf("Expected %v results, got: %v", tt.expectedResults, len(resp.Results))
			}

		})
	}

}

func TestUpdatePagination(t *testing.T) {
	cases := []struct {
		name             string
		config           *Config
		expectedErr      bool
		expectedNext     string
		expectedPrevious string
	}{
		{
			name:             "first page",
			config:           &Config{},
			expectedErr:      false,
			expectedNext:     baseURL + "/location-area/?offset=20&limit=20",
			expectedPrevious: "",
		},
		{
			name:             "second page",
			config:           &Config{},
			expectedErr:      false,
			expectedNext:     baseURL + "/location-area/?offset=40&limit=20",
			expectedPrevious: "/location-area/?offset=0&limit=20",
		},
		{
			name:             "missing url",
			config:           &Config{},
			expectedErr:      false,
			expectedNext:     baseURL + "/location-area/?offset=20&limit=20",
			expectedPrevious: "",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			resp := &NamedAPIResourceList{Next: tt.expectedNext, Previous: tt.expectedPrevious}
			err := tt.config.UpdatePagination(resp)

			if (err != nil) != tt.expectedErr {
				t.Logf("Expected err: %v but got: %v", tt.expectedErr, err)
			}

			if tt.config.Next.String() != tt.expectedNext {
				t.Errorf("Expected config Next: %v, actual: %v", tt.expectedNext, tt.config.Next.String())
			}

			if tt.config.Previous.String() != tt.expectedPrevious {
				t.Errorf("Expected config Previous: %v, actual: %v", tt.expectedPrevious, tt.config.Previous.String())
			}
		})
	}
}

func TestGetLocationAreas(t *testing.T) {

}
