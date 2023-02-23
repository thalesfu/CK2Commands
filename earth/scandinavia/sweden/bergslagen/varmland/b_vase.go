package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦瑟VaseBarony struct {
	feud.BaseBarony
}

var BVase韦瑟 feud.Barony = &韦瑟VaseBarony{}

func init() {
	f := BVase韦瑟.(*韦瑟VaseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vase",
		TitleName: "韦瑟",
		TitleCode: "b_vase",
	}
}
