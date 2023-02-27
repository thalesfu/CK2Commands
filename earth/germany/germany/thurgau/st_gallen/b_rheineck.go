package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赖内克RheineckBarony struct {
	feud.BaseBarony
}

var BRheineck赖内克 feud.Barony = &赖内克RheineckBarony{}

func init() {
    f := BRheineck赖内克.(*赖内克RheineckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rheineck",
		TitleName: "赖内克",
		TitleCode: "b_rheineck",
	}
}
