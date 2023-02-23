package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗德兹RodezBarony struct {
	feud.BaseBarony
}

var BRodez罗德兹 feud.Barony = &罗德兹RodezBarony{}

func init() {
	f := BRodez罗德兹.(*罗德兹RodezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rodez",
		TitleName: "罗德兹",
		TitleCode: "b_rodez",
	}
}
