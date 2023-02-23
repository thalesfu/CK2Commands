package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯瓦特波雷SvarteborgBarony struct {
	feud.BaseBarony
}

var BSvarteborg斯瓦特波雷 feud.Barony = &斯瓦特波雷SvarteborgBarony{}

func init() {
	f := BSvarteborg斯瓦特波雷.(*斯瓦特波雷SvarteborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svarteborg",
		TitleName: "斯瓦特波雷",
		TitleCode: "b_svarteborg",
	}
}
