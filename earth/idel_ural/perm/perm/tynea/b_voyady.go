package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃亚德VoyadyBarony struct {
	feud.BaseBarony
}

var BVoyady沃亚德 feud.Barony = &沃亚德VoyadyBarony{}

func init() {
    f := BVoyady沃亚德.(*沃亚德VoyadyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voyady",
		TitleName: "沃亚德",
		TitleCode: "b_voyady",
	}
}
