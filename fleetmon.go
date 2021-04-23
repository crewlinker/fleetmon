package fleetmon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	userAgent = "github.com/crewlinker/fleetmon:v0.0.1"
)

var defaultBaseURL = "https://apiv2.fleetmon.com/"

type VesselService service

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL // baseURL should always be specified with a trailing slash.
	apiKey     string
	userAgent  string

	Vessel *VesselService
}

type service struct {
	client *Client
}

func NewClient(apiKey string, httpClients ...*http.Client) *Client {
	var httpClient *http.Client

	if len(httpClients) > 0 && httpClients[0] != nil {
		httpClient = httpClients[0]
	} else {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    baseURL,
		userAgent:  userAgent,
	}
	s := &service{client: c}

	c.Vessel = (*VesselService)(s)

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(s string) (*http.Request, error) {
	rel, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Api-Key", c.apiKey)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}

	return resp, nil
}

func withOptions(s string, opt interface{}) (string, error) {
	if v := reflect.ValueOf(opt); v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
