package maritime

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

var baseURL = "https://services.marinetraffic.com"

// API is the interface for the maritime API.
type API interface {
	VesselMasterService // https://servicedocs.marinetraffic.com/tag/Search-Vessel
	VesselPhotoService  // https://servicedocs.marinetraffic.com/tag/Vessel-Information#operation/exportvesselphoto
	VesselSearchService // https://servicedocs.marinetraffic.com/tag/Search-Vessel

}

var _ API = (*Client)(nil)

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL // baseURL should always be specified with a trailing slash.
	apiKey     string
	userAgent  string
}

// NewClient returns a new maritime client.
func NewClient(apiKey string, httpClients ...*http.Client) *Client {
	var httpClient *http.Client

	if len(httpClients) > 0 && httpClients[0] != nil {
		httpClient = httpClients[0]
	} else {
		httpClient = http.DefaultClient
	}

	u, _ := url.Parse(baseURL)
	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURL:    u,
		userAgent:  "github.com/crewlinker/maritime:v0.0.1",
	}
}

// newRequest creates an API request.
func (c *Client) newRequest(ctx context.Context, path string, queryParams interface{}) (*http.Request, error) {
	u, err := url.Parse(fmt.Sprintf("/api/%s/%s", path, c.apiKey))
	if err != nil {
		return nil, err
	}

	values, err := query.Values(queryParams)
	if err != nil {
		return nil, err
	}

	u.RawQuery = values.Encode()

	reqUrl := c.baseURL.ResolveReference(u)

	fmt.Println(reqUrl.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var se ServiceError
		if err := json.NewDecoder(resp.Body).Decode(&se); err != nil {
			return nil, err
		}
		return nil, se
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, nil
}
