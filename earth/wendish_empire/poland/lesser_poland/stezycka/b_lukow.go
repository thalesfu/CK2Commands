package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 武库夫LukowBarony struct {
	feud.BaseBarony
}

var BLukow武库夫 feud.Barony = &武库夫LukowBarony{}

func init() {
    f := BLukow武库夫.(*武库夫LukowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lukow",
		TitleName: "武库夫",
		TitleCode: "b_lukow",
	}
}
