package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉赫RaaheBarony struct {
	feud.BaseBarony
}

var BRaahe拉赫 feud.Barony = &拉赫RaaheBarony{}

func init() {
	f := BRaahe拉赫.(*拉赫RaaheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raahe",
		TitleName: "拉赫",
		TitleCode: "b_raahe",
	}
}
