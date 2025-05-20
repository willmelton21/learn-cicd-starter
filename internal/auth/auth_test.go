package auth

import (
	"net/http"
	"testing"
	"reflect"
)


func TestAPIKey (t *testing.T) {
	
	tests := []struct {
		name string
		headers http.Header
		expectedKey string
		expectedError bool
	}{
		{
		name: "valid API Key",
		headers: http.Header{"Authorization": []string{"Bearer valid-api-key"}},
		expectedKey: "valid-api-key",
		expectedError: false,
	},
	{	name: "missing API key",
		headers: http.Header{},
		expectedKey: "",
		expectedError: true,
		},
	{	name: "malformed API Key",
		headers: http.Header{"Authorization": []string{"malformed-key-format"}},
		expectedKey: "",
		expectedError: true,
		},
	}

	for _, tc := range tests {
		got,err := GetAPIKey(tc.headers)
		if !reflect.DeepEqual(tc.expectedKey,got) || (tc.expectedError == true && err != nil) {
			t.Fatalf("test %s: expected: %v, got: %v, expected error: %v, got error: %v",tc.name,tc.expectedKey,got,tc.expectedError,err)	
		}
	}

}
