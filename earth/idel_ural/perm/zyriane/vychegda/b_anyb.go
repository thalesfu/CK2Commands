package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿内布AnybBarony struct {
	feud.BaseBarony
}

var BAnyb阿内布 feud.Barony = &阿内布AnybBarony{}

func init() {
    f := BAnyb阿内布.(*阿内布AnybBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anyb",
		TitleName: "阿内布",
		TitleCode: "b_anyb",
	}
}
