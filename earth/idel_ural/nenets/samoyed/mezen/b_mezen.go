package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅津MezenBarony struct {
	feud.BaseBarony
}

var BMezen梅津 feud.Barony = &梅津MezenBarony{}

func init() {
    f := BMezen梅津.(*梅津MezenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mezen",
		TitleName: "梅津",
		TitleCode: "b_mezen",
	}
}
