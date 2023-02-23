package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉韦洛RavelloBarony struct {
	feud.BaseBarony
}

var BRavello拉韦洛 feud.Barony = &拉韦洛RavelloBarony{}

func init() {
	f := BRavello拉韦洛.(*拉韦洛RavelloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ravello",
		TitleName: "拉韦洛",
		TitleCode: "b_ravello",
	}
}
