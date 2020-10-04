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
	url, err := parseURL("https://wanli.uno")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "wanli.uno" {
		t.Errorf("expected: %v, got: %v", "wanli.uno", url.Path)
	}

	_, err = parseURL("wanli.uno")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "wanli.uno"}
	if url.String() != "https://wanli.uno" {
		t.Errorf("expected: %v, got: %v", "https://wanli.uno", url.String())
	}

	url = URL{Scheme: "", Path: "wanli.uno"}
	if url.String() != "wanli.uno" {
		t.Errorf("expected: %v, got: %v", "wanli.uno", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "wanli.uno"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://wanli.uno\"" {
		t.Errorf("expected: %v, got: %v", "\"https://wanli.uno\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://wanli.uno\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "wanli.uno" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "wanli.uno"}, URL{"https", "wanli.uno"}, 0},
		{URL{"http", "wanli.uno"}, URL{"https", "wanli.uno"}, -1},
		{URL{"https", "wanli.uno/a"}, URL{"https", "wanli.uno"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "wanli.uno"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
