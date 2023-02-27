package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲拉赫VillachBarony struct {
	feud.BaseBarony
}

var BVillach菲拉赫 feud.Barony = &菲拉赫VillachBarony{}

func init() {
    f := BVillach菲拉赫.(*菲拉赫VillachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villach",
		TitleName: "菲拉赫",
		TitleCode: "b_villach",
	}
}
