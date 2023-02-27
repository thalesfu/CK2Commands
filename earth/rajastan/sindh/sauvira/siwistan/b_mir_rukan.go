package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弥尔楼乾Mir_rukanBarony struct {
	feud.BaseBarony
}

var BMir_rukan弥尔楼乾 feud.Barony = &弥尔楼乾Mir_rukanBarony{}

func init() {
    f := BMir_rukan弥尔楼乾.(*弥尔楼乾Mir_rukanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mir_rukan",
		TitleName: "弥尔楼乾",
		TitleCode: "b_mir_rukan",
	}
}
