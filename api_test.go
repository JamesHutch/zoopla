// Copyright 2014 James Hutchinson. All rights reserved.
//
// Use of this source code is governed by the Apache v2
// license that can be found in the LICENSE file.

package zoopla

import (
	"testing"
)

type TestOptions struct {
	Parameter1 string `url:"param1,omitempty"`
	Parameter2 string `url:"param2,omitempty"`
}

type NewRequestTest struct {
	api      *Api
	opts     TestOptions
	function string
	result   string
}

var nrTests = []NewRequestTest{
	{NewApi("key1"), TestOptions{Parameter1: "option1", Parameter2: "option2"}, "function1", "http://api.zoopla.co.uk/api/v1/function1?api_key=key1&param1=option1&param2=option2"},
	{NewApi("key2"), TestOptions{Parameter2: "option2"}, "function2", "http://api.zoopla.co.uk/api/v1/function2?api_key=key2&param2=option2"},
	{NewApi("key3"), TestOptions{Parameter1: "option1"}, "function3", "http://api.zoopla.co.uk/api/v1/function3?api_key=key3&param1=option1"},
}

func TestNewRequest(t *testing.T) {
	for _, test := range nrTests {
		res, err := test.api.NewRequest(test.function, "GET", test.opts)
		if err != nil {
			t.Error("NewRequest: unexpected error, ", err)
		}
		if res.URL.String() != test.result {
			t.Error("NewRequest: got " + res.URL.String() + " expected " + test.result)
		}
	}
}

type NewApiTest struct {
	key string
}

var newApiTests = []NewApiTest{
	{"key1"},
	{"key2"},
	{"key3"},
}

func TestNewApi(t *testing.T) {
	for _, test := range newApiTests {
		a := NewApi(test.key)
		if a.Key != test.key {
			t.Error("NewApi: got " + a.Key + " expected " + test.key)
		}
	}
}
