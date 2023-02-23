package caithness

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓比斯DunbeathBarony struct {
	feud.BaseBarony
}

var BDunbeath邓比斯 feud.Barony = &邓比斯DunbeathBarony{}

func init() {
	f := BDunbeath邓比斯.(*邓比斯DunbeathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunbeath",
		TitleName: "邓比斯",
		TitleCode: "b_dunbeath",
	}
}
