package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比列纳VillenaBarony struct {
	feud.BaseBarony
}

var BVillena比列纳 feud.Barony = &比列纳VillenaBarony{}

func init() {
    f := BVillena比列纳.(*比列纳VillenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villena",
		TitleName: "比列纳",
		TitleCode: "b_villena",
	}
}
