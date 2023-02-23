package temes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯韦瓦洛KevevaraBarony struct {
	feud.BaseBarony
}

var BKevevara凯韦瓦洛 feud.Barony = &凯韦瓦洛KevevaraBarony{}

func init() {
	f := BKevevara凯韦瓦洛.(*凯韦瓦洛KevevaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kevevara",
		TitleName: "凯韦瓦洛",
		TitleCode: "b_kevevara",
	}
}
