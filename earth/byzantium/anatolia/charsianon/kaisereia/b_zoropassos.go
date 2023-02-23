package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐罗帕索斯ZoropassosBarony struct {
	feud.BaseBarony
}

var BZoropassos佐罗帕索斯 feud.Barony = &佐罗帕索斯ZoropassosBarony{}

func init() {
	f := BZoropassos佐罗帕索斯.(*佐罗帕索斯ZoropassosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zoropassos",
		TitleName: "佐罗帕索斯",
		TitleCode: "b_zoropassos",
	}
}
