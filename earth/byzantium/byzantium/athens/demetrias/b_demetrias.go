package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德米特里阿斯DemetriasBarony struct {
	feud.BaseBarony
}

var BDemetrias德米特里阿斯 feud.Barony = &德米特里阿斯DemetriasBarony{}

func init() {
    f := BDemetrias德米特里阿斯.(*德米特里阿斯DemetriasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "demetrias",
		TitleName: "德米特里阿斯",
		TitleCode: "b_demetrias",
	}
}
