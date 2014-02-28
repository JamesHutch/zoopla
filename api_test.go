// Copyright 2014 James Hutchinson. All rights reserved.
//
// Use of this source code is governed by the Apache v2
// license that can be found in the LICENSE file.

package zoopla

import (
	"testing"
)

type TestOptions struct {
	parameter1 string
	parameter2 string
}

type RequestURLTest struct {
	api      Api
	opts     TestOptions
	function string
	result   string
}

func (o TestOptions) OptionString() string {
	return "parameter1=" + o.parameter1 + "&parameter2=" + o.parameter2
}

var urlTests = []RequestURLTest{
	{Api{"key1"}, TestOptions{"option1", "option2"}, "function1", "http://api.zoopla.co.uk/api/v1/function1.js?api_key=key1&parameter1=option1&parameter2=option2"},
	{Api{"世界"}, TestOptions{"option1", "option2"}, "function2", "http://api.zoopla.co.uk/api/v1/function2.js?api_key=世界&parameter1=option1&parameter2=option2"},
	{Api{"key1"}, TestOptions{"世界", "option2"}, "function3", "http://api.zoopla.co.uk/api/v1/function3.js?api_key=key1&parameter1=世界&parameter2=option2"},
}

func TestRequestURL(t *testing.T) {
	for _, test := range urlTests {
		res := test.api.RequestURL(test.function, test.opts)
		if res != test.result {
			t.Error("RequestURL: got " + res + " expected " + test.result)
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
		if a.key != test.key {
			t.Error("NewApi: got " + a.key + " expected " + test.key)
		}
	}
}
