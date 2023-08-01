package ipify

import (
	"net/http"
	"testing"
)

func TestGetIp(t *testing.T) {
	originalApiUri := API_URI_TEMPLATE

	// Create a dummy http.Request object
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	ip, err := GetIp(req)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Public IP address: %s", ip)
	}

	API_URI_TEMPLATE = "https://api.ipifyyyyyyyyyyyy.org?format=text&ip=%s"

	ip, err = GetIp(req)
	if err == nil {
		t.Error("Request to https://api.ipifyyyyyyyyyyyy.org should have failed, but succeeded.")
	}

	API_URI_TEMPLATE = originalApiUri
}
