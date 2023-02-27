package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆内拉MuneraBarony struct {
	feud.BaseBarony
}

var BMunera穆内拉 feud.Barony = &穆内拉MuneraBarony{}

func init() {
    f := BMunera穆内拉.(*穆内拉MuneraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munera",
		TitleName: "穆内拉",
		TitleCode: "b_munera",
	}
}
