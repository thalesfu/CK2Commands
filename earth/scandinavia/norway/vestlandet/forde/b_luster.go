package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕斯特LusterBarony struct {
	feud.BaseBarony
}

var BLuster吕斯特 feud.Barony = &吕斯特LusterBarony{}

func init() {
    f := BLuster吕斯特.(*吕斯特LusterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luster",
		TitleName: "吕斯特",
		TitleCode: "b_luster",
	}
}
