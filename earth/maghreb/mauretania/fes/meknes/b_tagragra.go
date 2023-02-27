package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔格拉格拉TagragraBarony struct {
	feud.BaseBarony
}

var BTagragra塔格拉格拉 feud.Barony = &塔格拉格拉TagragraBarony{}

func init() {
    f := BTagragra塔格拉格拉.(*塔格拉格拉TagragraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagragra",
		TitleName: "塔格拉格拉",
		TitleCode: "b_tagragra",
	}
}
