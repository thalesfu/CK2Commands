package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱杜尔加LedurgaBarony struct {
	feud.BaseBarony
}

var BLedurga莱杜尔加 feud.Barony = &莱杜尔加LedurgaBarony{}

func init() {
    f := BLedurga莱杜尔加.(*莱杜尔加LedurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ledurga",
		TitleName: "莱杜尔加",
		TitleCode: "b_ledurga",
	}
}
