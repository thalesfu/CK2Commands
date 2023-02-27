package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫兹多克MozdokBarony struct {
	feud.BaseBarony
}

var BMozdok莫兹多克 feud.Barony = &莫兹多克MozdokBarony{}

func init() {
    f := BMozdok莫兹多克.(*莫兹多克MozdokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mozdok",
		TitleName: "莫兹多克",
		TitleCode: "b_mozdok",
	}
}
