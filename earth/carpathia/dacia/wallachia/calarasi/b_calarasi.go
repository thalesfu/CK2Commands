package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克勒拉希CalarasiBarony struct {
	feud.BaseBarony
}

var BCalarasi克勒拉希 feud.Barony = &克勒拉希CalarasiBarony{}

func init() {
	f := BCalarasi克勒拉希.(*克勒拉希CalarasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calarasi",
		TitleName: "克勒拉希",
		TitleCode: "b_calarasi",
	}
}
