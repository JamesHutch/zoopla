// Copyright 2014 James Hutchinson. All rights reserved.
//
// Use of this source code is governed by the Apache v2
// license that can be found in the LICENSE file.

package zoopla

const baseURL = "http://api.zoopla.co.uk/api/v1/"

type Api struct {
	key string
}

func NewApi(key string) *Api {
	f := Api{key}
	return &f
}