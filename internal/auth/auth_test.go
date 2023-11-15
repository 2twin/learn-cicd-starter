package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	gotVal, gotErr := GetAPIKey(http.Header{})
	wantVal, wantErr := "", ErrNoAuthHeaderIncluded
	
	if gotVal != wantVal && gotErr != wantErr {
		t.Fatalf("expected value: %v, got value: %v;\n expected error: %v, got error: %v", wantVal, gotVal, wantErr, gotErr)
	}
}