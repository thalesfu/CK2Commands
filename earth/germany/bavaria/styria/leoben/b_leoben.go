package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱奥本LeobenBarony struct {
	feud.BaseBarony
}

var BLeoben莱奥本 feud.Barony = &莱奥本LeobenBarony{}

func init() {
    f := BLeoben莱奥本.(*莱奥本LeobenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leoben",
		TitleName: "莱奥本",
		TitleCode: "b_leoben",
	}
}
