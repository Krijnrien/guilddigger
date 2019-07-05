package gw2api

import "net/url"

// SkinDetails Either WeightClass(Armor) or DamageType(Weapon) is set depending
// on Type
type SkinDetails struct {
	Type        string `json:"type"`
	WeightClass string `json:"weight_class"`
	DamageType  string `json:"damage_type"`
}

// Skin all information on the skin
type Skin struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Flags        []string    `json:"flags"`
	Restrictions []string    `json:"restrictions"`
	Icon         string      `json:"icon"`
	Description  string      `json:"description"`
	Details      SkinDetails `json:"details"`
}

// Skins returns a list of all current skin ids
func (gw2 *GW2Api) Skins() (res []uint32, err error) {
	ver := "v2"
	tag := "skins"
	err = gw2.fetchEndpoint(ver, tag, nil, &res)
	return
}

// SkinIds returns the skin details as requested by the ids parameter.
func (gw2 *GW2Api) SkinIds(lang string, ids ...uint32) (skins []Skin, err error) {
	ver := "v2"
	tag := "skins"
	params := url.Values{}
	if lang != "" {
		params.Add("lang", lang)
	}
	params.Add("ids", commaListUint32(ids))
	err = gw2.fetchEndpoint(ver, tag, params, &skins)
	return
}
