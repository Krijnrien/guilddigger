package gw2api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CommerceListings returns a list of all current transactions ids
func (gw2 *GW2Api) CommerceListings() (res []uint32, err error) {
	ver := "v2"
	tag := "commerce/listings"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// CommerceListingIds returns the article listings for the provided ids
func (gw2 *GW2Api) CommerceListingIds(ids ...uint32) (articles []Listings, err error) {
	ver := "v2"
	tag := "commerce/listings"
	params := url.Values{}
	params.Add("ids", commaListUint32(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &articles)
	return
}

// CommerceListingPages for paginating through all existing listings
func (gw2 *GW2Api) CommerceListingPages(page int, pageSize int) (res []Listings, err error) {
	if page < 0 {
		return nil, fmt.Errorf("page parameter cannot be a negative number")
	}

	ver := "v2"
	tag := "commerce/listings"
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	if pageSize >= 0 {
		params.Add("page_size", strconv.Itoa(pageSize))
	}
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

// Exchange with CoinsPerGem and Quantity. Varies on request
type Exchange struct {
	CoinsPerGem int `json:"coins_per_gem"`
	Quantity    int `json:"quantity"`
}

// CommerceExchangeGems returns the amount of gold given for the quantity of
// coin
func (gw2 *GW2Api) CommerceExchangeGems(quantity int) (res Exchange, err error) {
	if quantity < 1 {
		return res, fmt.Errorf("required parameter too low")
	}

	ver := "v2"
	tag := "commerce/exchange/gems"
	params := url.Values{}
	params.Add("quantity", strconv.Itoa(quantity))
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

// CommerceExchangeCoins returns the amount of gems given for the quantity of
// gems
func (gw2 *GW2Api) CommerceExchangeCoins(quantity int64) (res Exchange, err error) {
	if quantity < 1 {
		return res, fmt.Errorf("required parameter too low")
	}

	ver := "v2"
	tag := "commerce/exchange/coins"
	params := url.Values{}
	params.Add("quantity", strconv.FormatInt(quantity, 10))
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

func (gw2 *GW2Api) CommercePricePageHeaders(page int, pageSize int) (body *http.Response, err error) {
	ver := "v2"
	tag := "commerce/prices"
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("page_size", strconv.Itoa(pageSize))

	var endpoint *url.URL
	endpoint, _ = url.Parse("https://api.guildwars2.com")
	endpoint.Path += "/" + ver + "/" + tag
	endpoint.RawQuery = params.Encode()

	// not checking for params != nil since its manually set in this function.
	body, err = gw2.fetchRawEndpointHeaders(endpoint.String())
	return
}

// CommercePrices returns a list of all ids
func (gw2 *GW2Api) CommercePrices() (res []uint32, err error) {
	ver := "v2"
	tag := "commerce/prices"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// CommercePriceIds returns price information about the requested ids
func (gw2 *GW2Api) CommercePriceIds(ids ...uint32) (artprices []Prices, err error) {
	ver := "v2"
	tag := "commerce/prices"
	params := url.Values{}
	params.Add("ids", commaListUint32(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &artprices)
	return
}

// CommerceListingPages for paginating through all existing listings
func (gw2 *GW2Api) CommercePricePages(page uint32, pageSize uint32) (res []Prices, err error) {
	ver := "v2"
	tag := "commerce/prices"
	params := url.Values{}
	params.Add("page", strconv.FormatUint(uint64(page), 10))
	params.Add("page_size", strconv.FormatUint(uint64(pageSize), 10))
	err = gw2.fetchEndpoint(ver, tag, params, &res)
	return
}

// Transaction represents one of the accounts listed buy or sell orders
type Transaction struct {
	ID        int    `json:"id"`
	ItemID    int    `json:"item_id"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Created   string `json:"created"`
	Purchased string `json:"purchased"`
}

// CommerceTransactionsCurrentBuys returns all current buy orders of the account
func (gw2 *GW2Api) CommerceTransactionsCurrentBuys() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/buys"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

// CommerceTransactionsCurrentSells returns all current sell orders of the account
func (gw2 *GW2Api) CommerceTransactionsCurrentSells() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/sells"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

// CommerceTransactionsHistoryBuys returns all past buy orders of the account
// for the last 90 days
func (gw2 *GW2Api) CommerceTransactionsHistoryBuys() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/current/buys"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}

// CommerceTransactionsHistorySells returns all past sell orders of the account
// for the last 90 days
func (gw2 *GW2Api) CommerceTransactionsHistorySells() (trans []Transaction, err error) {
	ver := "v2"
	tag := "commerce/transactions/history/sells"
	err = gw2.fetchAuthenticatedEndpoint(ver, tag, PermTradingpost, nil, &trans)
	return
}
