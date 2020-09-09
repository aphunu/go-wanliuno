// Copyright 2018 The go-wanliuno Authors
// This file is part of the go-wanliuno library.
//
// The go-wanliuno library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-wanliuno library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-wanliuno library. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"testing"
)

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://wanliuno.org")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "wanliuno.org" {
		t.Errorf("expected: %v, got: %v", "wanliuno.org", url.Path)
	}

	_, err = parseURL("wanliuno.org")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "wanliuno.org"}
	if url.String() != "https://wanliuno.org" {
		t.Errorf("expected: %v, got: %v", "https://wanliuno.org", url.String())
	}

	url = URL{Scheme: "", Path: "wanliuno.org"}
	if url.String() != "wanliuno.org" {
		t.Errorf("expected: %v, got: %v", "wanliuno.org", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "wanliuno.org"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://wanliuno.org\"" {
		t.Errorf("expected: %v, got: %v", "\"https://wanliuno.org\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://wanliuno.org\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "wanliuno.org" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "wanliuno.org"}, URL{"https", "wanliuno.org"}, 0},
		{URL{"http", "wanliuno.org"}, URL{"https", "wanliuno.org"}, -1},
		{URL{"https", "wanliuno.org/a"}, URL{"https", "wanliuno.org"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "wanliuno.org"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
