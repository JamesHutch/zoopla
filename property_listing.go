// Copyright 2014 James Hutchinson. All rights reserved.
//
// Use of this source code is governed by the Apache v2
// license that can be found in the LICENSE file.

// Documentation on http://developer.zoopla.com/docs/read/Property_listings

package zoopla

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/JamesHutch/recode"
)

const (
	apiURL    = "property_listings.js"
	apiMethod = "GET"
)

type PropertyListingOptions struct {
	Area          string  `url:"area,omitempty" json:"area,omitempty"` // Arbitrary area name, or postcode.
	Street        string  `url:"street,omitempty" json:"street,omitempty"`
	Town          string  `url:"town,omitempty" json:"town,omitempty"`
	Postcode      string  `url:"postcode,omitempty" json:"postcode,omitempty"`
	County        string  `url:"county,omitempty" json:"county,omitempty"`
	Country       string  `url:"country,omitempty" json:"country,omitempty"`
	Latitude      float64 `url:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude     float64 `url:"longitude,omitempty" json:"longitude,omitempty"`
	LatMin        float64 `url:"lat_min,omitempty" json:"lat_min,omitempty"`
	LatMax        float64 `url:"lat_max,omitempty" json:"lat_max,omitempty"`
	LonMin        float64 `url:"lon_min,omitempty" json:"lon_min,omitempty"`
	LonMax        float64 `url:"lon_max,omitempty" json:"lon_max,omitempty"`
	OutputType    string  `url:"output_type,omitempty" json:"output_type,omitempty"`       // The actual area type to restrict the location request to - postcode, outcode, street, town, area or county. For instance, specifying a value for "postcode" (or a postcode within the "area" parameter) above and then a value of "town" for this parameter will use the town that the postcode is within, not the postcode itself. Note that the same occurs for latitude/longitude searches; providing an exact location with an output type of "postcode" will use the postcode during the request, not the latitude/longitude provided.
	Radius        float64 `url:"radius,omitempty" json:"radius,omitempty"`                 // From 0.5 to 40
	OrderBy       string  `url:"order_by,omitempty" json:"order_by,omitempty"`             // "price" (default) or "age"
	Ordering      string  `url:"ordering,omitempty" json:"ordering,omitempty"`             // "descending" (default) or "ascending"
	ListingStatus string  `url:"listing_status,omitempty" json:"listing_status,omitempty"` // "sale" or "rent"
	IncludeSold   string  `url:"include_sold,omitempty" json:"include_sold,omitempty"`     // "1" or "0". Defaults to 0.
	IncludeRented string  `url:"include_rented,omitempty" json:"include_rented,omitempty"` // "1" or "0". Defaults to 0.
	MinimumPrice  int     `url:"minimum_price,omitempty" json:"minimum_price,omitempty"`   //  When listing_status is "sale" this refers to the sale price and when listing_status is "rent" it refers to the per-week price.
	MaximumPrice  int     `url:"maximum_price,omitempty" json:"maximum_price,omitempty"`
	MinimumBeds   int     `url:"minimum_beds,omitempty" json:"minimum_beds,omitempty"`
	MaximumBeds   int     `url:"maximum_beds,omitempty" json:"maximum_beds,omitempty"`
	Furnished     string  `url:"furnished,omitempty" json:"furnished,omitempty"`         // "furnished", "unfurnished" or "part-furnished"
	PropertyType  string  `url:"property_type,omitempty" json:"property_type,omitempty"` // "houses" or "flats"
	NewHomes      string  `url:"new_homes,omitempty" json:"new_homes,omitempty"`         // Specifying "yes"/"true" will restrict to only new homes, "no"/"false" will exclude them from the results set.
	ChainFree     string  `url:"chain_free,omitempty" json:"chain_free,omitempty"`       // Specifying "yes"/"true" will restrict to chain free homes, "no"/"false" will exclude them from the results set.
	Keywords      string  `url:"keywords,omitempty" json:"keywords,omitempty"`
	ListingID     string  `url:"listing_id,omitempty" json:"listing_id,omitempty"`
	BranchID      string  `url:"branch_id,omitempty" json:"branch_id,omitempty"`
	PageNumber    int     `url:"page_number,omitempty" json:"page_number,omitempty"` // default 1
	PageSize      int     `url:"page_size,omitempty" json:"page_size,omitempty"`     // default 10, maximum 100
	Summarised    string  `url:"summarised,omitempty" json:"summarised,omitempty"`   // Specifying "yes"/"true" will return a cut-down entry for each listing with the description cut short and the following fields will be removed: price_change, floor_plan.
}

type RawPriceChange struct {
	Price interface{} `json:"price,omitempty"`
	Date  string      `json:"date,omitempty"`
}

type RawProperyListing struct {
	ListingID          string           `json:"listing_id"`
	Outcode            string           `json:"outcode,omitempty"`
	DisplayableAddress string           `json:"displayable_address,omitempty"`
	County             string           `json:"county,omitempty"`
	Country            string           `json:"country,omitempty"`
	NumBathrooms       interface{}      `json:"num_bathrooms,omitempty"`
	NumBedrooms        interface{}      `json:"num_bedrooms,omitempty"`
	NumFloors          interface{}      `json:"num_floors,omitempty"`
	NumRecepts         interface{}      `json:"num_recepts,omitempty"`
	ListingStatus      string           `json:"listing_status,omitempty"`
	Status             string           `json:"status,omitempty"`
	Price              interface{}      `json:"price,omitempty"`
	PriceModifier      string           `json:"price_modifier,omitempty"`
	PriceChange        []RawPriceChange `json:"price_change,omitempty"`
	PropertyType       string           `json:"property_type,omitempty"`
	StreetName         string           `json:"street_name,omitempty"`
	ThumbnailURL       string           `json:"thumbnail_url,omitempty"`
	ImageURL           string           `json:"image_url,omitempty"`
	ImageCaption       string           `json:"image_caption,omitempty"`
	FloorPlan          []string         `json:"floor_plan,omitempty"`
	Description        string           `json:"description,omitempty"`
	ShortDescription   string           `json:"short_description,omitempty"`
	DetailsURL         string           `json:"details_url,omitempty"`
	NewHome            string           `json:"new_home,omitempty"`
	Latitude           float64          `json:"latitude,omitempty"`
	Longitude          float64          `json:"longitude,omitempty"`
	FirstPublishedDate string           `json:"first_published_date,omitempty"`
	LastPublishedDate  string           `json:"last_published_date,omitempty"`
	AgentName          string           `json:"agent_name,omitempty"`
	AgentLogo          string           `json:"agent_logo,omitempty"`
	AgentPhone         string           `json:"agent_phone,omitempty"`
}

type RawPropertyListingResults struct {
	ResultCount int                 `json:"result_count"`
	Listings    []RawProperyListing `json:"listing"`
}

type PriceChange struct {
	Price float64
	Date  string
}

type ProperyListing struct {
	ListingID          string        `json:"listing_id"`
	Outcode            string        `json:"outcode,omitempty"`
	DisplayableAddress string        `json:"displayable_address,omitempty"`
	County             string        `json:"county,omitempty"`
	Country            string        `json:"country,omitempty"`
	NumBathrooms       float64       `json:"num_bathrooms,omitempty"`
	NumBedrooms        float64       `json:"num_bedrooms,omitempty"`
	NumFloors          float64       `json:"num_floors,omitempty"`
	NumRecepts         float64       `json:"num_recepts,omitempty"`
	ListingStatus      string        `json:"listing_status,omitempty"`
	Status             string        `json:"status,omitempty"`
	Price              float64       `json:"price,omitempty"`
	PriceModifier      string        `json:"price_modifier,omitempty"`
	PriceChange        []PriceChange `json:"price_change,omitempty"`
	PropertyType       string        `json:"property_type,omitempty"`
	StreetName         string        `json:"street_name,omitempty"`
	ThumbnailURL       string        `json:"thumbnail_url,omitempty"`
	ImageURL           string        `json:"image_url,omitempty"`
	ImageCaption       string        `json:"image_caption,omitempty"`
	FloorPlan          []string      `json:"floor_plan,omitempty"`
	Description        string        `json:"description,omitempty"`
	ShortDescription   string        `json:"short_description,omitempty"`
	DetailsURL         string        `json:"details_url,omitempty"`
	NewHome            string        `json:"new_home,omitempty"`
	Latitude           float64       `json:"latitude,omitempty"`
	Longitude          float64       `json:"longitude,omitempty"`
	FirstPublishedDate string        `json:"first_published_date,omitempty"`
	LastPublishedDate  string        `json:"last_published_date,omitempty"`
	AgentName          string        `json:"agent_name,omitempty"`
	AgentLogo          string        `json:"agent_logo,omitempty"`
	AgentPhone         string        `json:"agent_phone,omitempty"`
}

type PropertyListingResults struct {
	ResultCount int              `json:"result_count"`
	Listings    []ProperyListing `json:"listing"`
}

func (a *Api) GetListings(opt *PropertyListingOptions) (*PropertyListingResults, error) {
	req, err := a.NewRequest(apiURL, apiMethod, opt)
	if err != nil {
		return nil, err
	}
	log.Print(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rawResult RawPropertyListingResults
	err = json.Unmarshal(body, &rawResult)
	if err != nil {
		return nil, err
	}
	var result PropertyListingResults
	err = recode.Recode(rawResult, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *Api) GetListingsWithValues(v url.Values) (*PropertyListingResults, error) {
	req, err := a.NewRequest(apiURL, apiMethod, PropertyListingOptions{})
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for key, values := range v {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()
	log.Printf("out request: %v", req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rawResult RawPropertyListingResults
	err = json.Unmarshal(body, &rawResult)
	if err != nil {
		return nil, err
	}
	var result PropertyListingResults
	err = recode.Recode(rawResult, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
