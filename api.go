// Copyright 2014 James Hutchinson. All rights reserved.
//
// Use of this source code is governed by the Apache v2
// license that can be found in the LICENSE file.

package zoopla

import (
    "net/http"
    "net/url"

    "github.com/google/go-querystring/query"
)

const (
    version        = "0.01"
    defaultBaseURL = "http://api.zoopla.co.uk/api/v1/"
    userAgent      = "github.com/JamesHutch/zoopla/" + version
)

type Api struct {
    BaseURL   *url.URL
    Key       string
    UserAgent string
}

func NewApi(key string) *Api {
    baseURL, _ := url.Parse(defaultBaseURL)
    return &Api{Key: key, UserAgent: userAgent, BaseURL: baseURL}
}

func (a *Api) NewRequest(relUrl, method string, opt interface{}) (*http.Request, error) {
    rel, err := url.Parse(relUrl)
    if err != nil {
        return nil, err
    }
    u := a.BaseURL.ResolveReference(rel)
    qs, err := query.Values(opt)
    if err != nil {
        return nil, err
    }
    qs.Set("api_key", a.Key)
    u.RawQuery = qs.Encode()
    req, err := http.NewRequest(method, u.String(), nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("User-Agent", a.UserAgent)
    return req, nil
}
