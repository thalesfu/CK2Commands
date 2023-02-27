package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科洛尔诺ColornoBarony struct {
	feud.BaseBarony
}

var BColorno科洛尔诺 feud.Barony = &科洛尔诺ColornoBarony{}

func init() {
    f := BColorno科洛尔诺.(*科洛尔诺ColornoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "colorno",
		TitleName: "科洛尔诺",
		TitleCode: "b_colorno",
	}
}
