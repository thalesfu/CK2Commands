package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索洛图恩SolothurnBarony struct {
	feud.BaseBarony
}

var BSolothurn索洛图恩 feud.Barony = &索洛图恩SolothurnBarony{}

func init() {
    f := BSolothurn索洛图恩.(*索洛图恩SolothurnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "solothurn",
		TitleName: "索洛图恩",
		TitleCode: "b_solothurn",
	}
}
