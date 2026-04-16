package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAsciiArtPageHandler(t *testing.T) {
	loadDefaults()

	type record struct {
		form     url.Values
		wantCode int
	}
	test := record{}
	tests := []record{}

	// 1st test - normal input
	form := url.Values{}
	form.Add("text", "ABCD\nEFGH")
	form.Add("color", "#ffffff")
	form.Add("colored-text", "")
	form.Add("alignment", "left")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusFound
	tests = append(tests, test)

	// 2nd test - bad banner
	form = url.Values{}
	form.Add("text", "ABCD\nEFGH")
	form.Add("color", "#ffffff")
	form.Add("colored-text", "")
	form.Add("alignment", "left")
	form.Add("banner", "bad banner")
	test.form = form
	test.wantCode = http.StatusNotFound
	tests = append(tests, test)

	// 3rd test - empty text
	form = url.Values{}
	form.Add("text", "")
	form.Add("color", "#ffffff")
	form.Add("colored-text", "")
	form.Add("alignment", "left")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusBadRequest
	tests = append(tests, test)

	// 4th test - bad color
	form = url.Values{}
	form.Add("text", "ABCD\nEFGH")
	form.Add("color", "gold")
	form.Add("colored-text", "")
	form.Add("alignment", "left")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusBadRequest
	tests = append(tests, test)

	// 5th test - non-ASCII colored-text
	form = url.Values{}
	form.Add("text", "ABCD\nEFGH")
	form.Add("color", "#ffffff")
	form.Add("colored-text", "❌")
	form.Add("alignment", "left")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusFound
	tests = append(tests, test)

	// 6th test - bad alignment
	form = url.Values{}
	form.Add("text", "ABCD\nEFGH")
	form.Add("color", "#ffffff")
	form.Add("colored-text", "")
	form.Add("alignment", "bad alignment")
	form.Add("banner", "standard")
	test.form = form
	test.wantCode = http.StatusBadRequest
	tests = append(tests, test)

	// 7th test - non-ASCII text
	form = url.Values{}
	form.Add("text", "❌")
	form.Add("color", "#ffffff")
	form.Add("colored-text", "")
	form.Add("alignment", "left")
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
			t.Errorf("expected %d got %d", test.wantCode, result.StatusCode)
		}
	}

}
