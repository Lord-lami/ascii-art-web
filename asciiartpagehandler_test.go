package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAsciiArtPageHandler(t *testing.T) {
	loadBanners()

	type record struct {
		form     url.Values
		wantCode int
	}
	test := record{}
	tests := []record{}

	// 1st test - normal input
	form := url.Values{}
	form.Add("text", "ABCD\nEFGH")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusFound
	tests = append(tests, test)

	// 2nd test - bad banner
	form = url.Values{}
	form.Add("text", "ABCD\nEFGH")
	form.Add("banner", "bad banner")
	test.form = form
	test.wantCode = http.StatusBadRequest
	tests = append(tests, test)

	// 3rd test - empty text
	form = url.Values{}
	form.Add("text", "")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusBadRequest
	tests = append(tests, test)

	// 4th test - non-ASCII text
	form = url.Values{}
	form.Add("text", "❌")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusInternalServerError
	tests = append(tests, test)

	for _, test := range tests {
		req := httptest.NewRequest("POST", "/ascii-art", strings.NewReader(test.form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		asciiArtPageHandler(w, req)
		result := w.Result()
		// body, _ := io.ReadAll(result.Body)

		if result.StatusCode != test.wantCode {
			t.Errorf("expected 301 got %d", result.StatusCode)
		}
	}

}
