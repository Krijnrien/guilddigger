package gw2api

import (
	"fmt"
	"github.com/krijnrien/microguild/pkg/messages"
	"net/url"
	"strconv"
)

// Items returns a list of all item ids.
func (gw2 *GW2Api) Items() (res []uint32, err error) {
	ver := "v2"
	tag := "items"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// ItemDetails returns a list of items depending on the parameters either with
// pagination or as requested by ids
func (gw2 *GW2Api) ItemDetails(page int, pageSize int, lang string, ids ...uint32) ([]messages.Item, error) {
	if ids != nil {
		return gw2.ItemIds(lang, ids...)
	} else if page >= 0 {
		return gw2.ItemPages(page, pageSize, lang)
	} else {
		return nil, fmt.Errorf("invalid combination of parameters. Consider using Items() instead")
	}
}

// ItemIds returns a list of detailed item description for the requested ids.
// Names of the items are localized according to the lang parameter
func (gw2 *GW2Api) ItemIds(lang string, ids ...uint32) (items []messages.Item, err error) {
	ver := "v2"
	tag := "items"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaListUint32(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &items)
	return
}

// ItemPages returns a list of detailed item description for the requested page.
// Names of the items are localized according to the lang parameter
func (gw2 *GW2Api) ItemPages(page int, pageSize int, lang string) (items []messages.Item, err error) {
	if page < 0 {
		return nil, fmt.Errorf("page parameter cannot be a negative number")
	}

	ver := "v2"
	tag := "items"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("page", strconv.Itoa(page))
	if pageSize >= 0 {
		params.Add("page_size", strconv.Itoa(pageSize))
	}
	err = gw2.fetchEndpoint(ver, tag, params, &items)
	return
}
