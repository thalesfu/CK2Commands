package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡特梅尔CartmelBarony struct {
	feud.BaseBarony
}

var BCartmel卡特梅尔 feud.Barony = &卡特梅尔CartmelBarony{}

func init() {
	f := BCartmel卡特梅尔.(*卡特梅尔CartmelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cartmel",
		TitleName: "卡特梅尔",
		TitleCode: "b_cartmel",
	}
}
