package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达里耶巴德DariyabadBarony struct {
	feud.BaseBarony
}

var BDariyabad达里耶巴德 feud.Barony = &达里耶巴德DariyabadBarony{}

func init() {
	f := BDariyabad达里耶巴德.(*达里耶巴德DariyabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dariyabad",
		TitleName: "达里耶巴德",
		TitleCode: "b_dariyabad",
	}
}
