package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉赫LaachBarony struct {
	feud.BaseBarony
}

var BLaach拉赫 feud.Barony = &拉赫LaachBarony{}

func init() {
	f := BLaach拉赫.(*拉赫LaachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laach",
		TitleName: "拉赫",
		TitleCode: "b_laach",
	}
}
