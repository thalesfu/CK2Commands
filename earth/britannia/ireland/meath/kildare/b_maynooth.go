package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅努斯MaynoothBarony struct {
	feud.BaseBarony
}

var BMaynooth梅努斯 feud.Barony = &梅努斯MaynoothBarony{}

func init() {
    f := BMaynooth梅努斯.(*梅努斯MaynoothBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maynooth",
		TitleName: "梅努斯",
		TitleCode: "b_maynooth",
	}
}
