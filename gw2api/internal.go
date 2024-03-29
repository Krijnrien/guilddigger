package gw2api

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

func (gw2 *GW2Api) dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, gw2.timeout)
}

func (gw2 *GW2Api) fetchRawEndpoint(url string) (io io.ReadCloser, err error) {
	var resp *http.Response
	if resp, err = gw2.Client.Get(url); err != nil {
		return
	}
	return resp.Body, nil
}

func (gw2 *GW2Api) fetchRawEndpointHeaders(url string) (resp *http.Response, err error) {
	if resp, err = gw2.Client.Get(url); err != nil {
		return
	}
	return
}

func (gw2 *GW2Api) fetchEndpoint(ver, tag string, params url.Values, result interface{}) (err error) {
	var endpoint *url.URL
	endpoint, _ = url.Parse("https://api.guildwars2.com")
	endpoint.Path += "/" + ver + "/" + tag
	if params != nil {
		endpoint.RawQuery = params.Encode()
	}

	var resp *http.Response
	if resp, err = gw2.Client.Get(endpoint.String()); err != nil {
		return err
	}

	err2 := json.NewDecoder(resp.Body).Decode(&result)
	if err2 != nil {
		return err2
	}
	defer resp.Body.Close()
	return
}

func (gw2 *GW2Api) fetchAuthenticatedEndpoint(ver, tag string, perm Permission, params url.Values, result interface{}) (err error) {
	if len(gw2.auth) < 1 {
		return fmt.Errorf("API requires authentication")
	}

	if perm >= PermAccount && !flagGet(gw2.authFlags, uint(perm)) {
		return fmt.Errorf("missing permissions for authenticated Endpoint")
	}

	if params == nil {
		params = url.Values{}
	}
	params.Add("access_token", gw2.auth)
	return gw2.fetchEndpoint(ver, tag, params, result)
}
