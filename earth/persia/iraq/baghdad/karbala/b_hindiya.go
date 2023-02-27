package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛蒂亚HindiyaBarony struct {
	feud.BaseBarony
}

var BHindiya辛蒂亚 feud.Barony = &辛蒂亚HindiyaBarony{}

func init() {
    f := BHindiya辛蒂亚.(*辛蒂亚HindiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hindiya",
		TitleName: "辛蒂亚",
		TitleCode: "b_hindiya",
	}
}
