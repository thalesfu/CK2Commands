package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍奇诺KhotchinoBarony struct {
	feud.BaseBarony
}

var BKhotchino霍奇诺 feud.Barony = &霍奇诺KhotchinoBarony{}

func init() {
	f := BKhotchino霍奇诺.(*霍奇诺KhotchinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khotchino",
		TitleName: "霍奇诺",
		TitleCode: "b_khotchino",
	}
}
