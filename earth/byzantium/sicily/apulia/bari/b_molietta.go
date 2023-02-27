package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫列塔MoliettaBarony struct {
	feud.BaseBarony
}

var BMolietta莫列塔 feud.Barony = &莫列塔MoliettaBarony{}

func init() {
    f := BMolietta莫列塔.(*莫列塔MoliettaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molietta",
		TitleName: "莫列塔",
		TitleCode: "b_molietta",
	}
}
