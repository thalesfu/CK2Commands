package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 津陶雷夫SzintarevBarony struct {
	feud.BaseBarony
}

var BSzintarev津陶雷夫 feud.Barony = &津陶雷夫SzintarevBarony{}

func init() {
	f := BSzintarev津陶雷夫.(*津陶雷夫SzintarevBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szintarev",
		TitleName: "津陶雷夫",
		TitleCode: "b_szintarev",
	}
}
