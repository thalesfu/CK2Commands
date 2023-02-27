package seljuk_turks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Seljuk_turksEmpire interface {
    feud.Empire
}

type 塞尔柱帝国Seljuk_turksEmpire struct {
	feud.BaseEmpire
}

var ESeljuk_turks塞尔柱帝国 Seljuk_turksEmpire = &塞尔柱帝国Seljuk_turksEmpire{}

func init() {
	f := ESeljuk_turks塞尔柱帝国.(*塞尔柱帝国Seljuk_turksEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "seljuk_turks",
		TitleName: "塞尔柱帝国",
		TitleCode: "e_seljuk_turks",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
