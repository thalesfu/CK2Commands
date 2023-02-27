package ikh_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湖谷Nuuruudyn_hondijBarony struct {
	feud.BaseBarony
}

var BNuuruudyn_hondij湖谷 feud.Barony = &湖谷Nuuruudyn_hondijBarony{}

func init() {
    f := BNuuruudyn_hondij湖谷.(*湖谷Nuuruudyn_hondijBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nuuruudyn_hondij",
		TitleName: "湖谷",
		TitleCode: "b_nuuruudyn_hondij",
	}
}
