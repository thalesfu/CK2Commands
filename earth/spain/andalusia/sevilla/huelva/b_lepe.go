package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱佩LepeBarony struct {
	feud.BaseBarony
}

var BLepe莱佩 feud.Barony = &莱佩LepeBarony{}

func init() {
	f := BLepe莱佩.(*莱佩LepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lepe",
		TitleName: "莱佩",
		TitleCode: "b_lepe",
	}
}
