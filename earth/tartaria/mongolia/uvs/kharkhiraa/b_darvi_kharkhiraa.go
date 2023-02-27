package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔维Darvi_kharkhiraaBarony struct {
	feud.BaseBarony
}

var BDarvi_kharkhiraa达尔维 feud.Barony = &达尔维Darvi_kharkhiraaBarony{}

func init() {
    f := BDarvi_kharkhiraa达尔维.(*达尔维Darvi_kharkhiraaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darvi_kharkhiraa",
		TitleName: "达尔维",
		TitleCode: "b_darvi_kharkhiraa",
	}
}
