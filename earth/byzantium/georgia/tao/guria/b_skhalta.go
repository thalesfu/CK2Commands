package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯哈尔塔SkhaltaBarony struct {
	feud.BaseBarony
}

var BSkhalta斯哈尔塔 feud.Barony = &斯哈尔塔SkhaltaBarony{}

func init() {
	f := BSkhalta斯哈尔塔.(*斯哈尔塔SkhaltaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skhalta",
		TitleName: "斯哈尔塔",
		TitleCode: "b_skhalta",
	}
}
