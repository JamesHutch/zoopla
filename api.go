// Copyright 2014 James Hutchinson. All rights reserved.
//
// Use of this source code is governed by the Apache v2
// license that can be found in the LICENSE file.

package zoopla

import (
	"io/ioutil"
	"net/http"
)

const BaseURL = "http://api.zoopla.co.uk/api/v1/"

type OptionStringer interface {
	OptionString() string
}

type Api struct {
	key string
}

func NewApi(key string) *Api {
	f := Api{key}
	return &f
}

func (a *Api) RequestURL(function string, opt OptionStringer) string {
	return BaseURL + function + ".js?api_key=" + a.key + "&" + opt.OptionString()
}
