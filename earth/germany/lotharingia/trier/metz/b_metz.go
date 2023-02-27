package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅斯MetzBarony struct {
	feud.BaseBarony
}

var BMetz梅斯 feud.Barony = &梅斯MetzBarony{}

func init() {
    f := BMetz梅斯.(*梅斯MetzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "metz",
		TitleName: "梅斯",
		TitleCode: "b_metz",
	}
}
