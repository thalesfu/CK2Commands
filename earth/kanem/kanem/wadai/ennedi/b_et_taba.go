package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔巴Et_tabaBarony struct {
	feud.BaseBarony
}

var BEt_taba塔巴 feud.Barony = &塔巴Et_tabaBarony{}

func init() {
    f := BEt_taba塔巴.(*塔巴Et_tabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "et_taba",
		TitleName: "塔巴",
		TitleCode: "b_et_taba",
	}
}
