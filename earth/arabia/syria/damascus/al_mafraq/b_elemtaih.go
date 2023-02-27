package al_mafraq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾勒穆塔ElemtaihBarony struct {
	feud.BaseBarony
}

var BElemtaih艾勒穆塔 feud.Barony = &艾勒穆塔ElemtaihBarony{}

func init() {
    f := BElemtaih艾勒穆塔.(*艾勒穆塔ElemtaihBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elemtaih",
		TitleName: "艾勒穆塔",
		TitleCode: "b_elemtaih",
	}
}
