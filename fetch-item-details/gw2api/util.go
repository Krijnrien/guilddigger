package gw2api

import (
	"fmt"
	"strings"
)

func commaListUint32(ids []uint32) string {
	return strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
}

func commaListInt(ids []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
}

func commaListString(ids []string) string {
	return strings.Join(ids, ",")
}

func flagGet(n, pos uint) bool {
	return (n>>pos)&1 == 1
}

func flagSet(n *uint, pos uint) {
	*n |= 1 << pos
}
