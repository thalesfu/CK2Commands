package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里瓦德奥RibadeoBarony struct {
	feud.BaseBarony
}

var BRibadeo里瓦德奥 feud.Barony = &里瓦德奥RibadeoBarony{}

func init() {
    f := BRibadeo里瓦德奥.(*里瓦德奥RibadeoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ribadeo",
		TitleName: "里瓦德奥",
		TitleCode: "b_ribadeo",
	}
}
