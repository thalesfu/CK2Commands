package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默德灵ModlingBarony struct {
	feud.BaseBarony
}

var BModling默德灵 feud.Barony = &默德灵ModlingBarony{}

func init() {
    f := BModling默德灵.(*默德灵ModlingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "modling",
		TitleName: "默德灵",
		TitleCode: "b_modling",
	}
}
