package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉努夫LanufBarony struct {
	feud.BaseBarony
}

var BLanuf拉努夫 feud.Barony = &拉努夫LanufBarony{}

func init() {
	f := BLanuf拉努夫.(*拉努夫LanufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lanuf",
		TitleName: "拉努夫",
		TitleCode: "b_lanuf",
	}
}
