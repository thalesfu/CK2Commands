package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱希特格明德LechtgemundBarony struct {
	feud.BaseBarony
}

var BLechtgemund莱希特格明德 feud.Barony = &莱希特格明德LechtgemundBarony{}

func init() {
    f := BLechtgemund莱希特格明德.(*莱希特格明德LechtgemundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lechtgemund",
		TitleName: "莱希特格明德",
		TitleCode: "b_lechtgemund",
	}
}
