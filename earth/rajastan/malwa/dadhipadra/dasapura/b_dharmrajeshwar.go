package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达摩罗逝湿伐罗DharmrajeshwarBarony struct {
	feud.BaseBarony
}

var BDharmrajeshwar达摩罗逝湿伐罗 feud.Barony = &达摩罗逝湿伐罗DharmrajeshwarBarony{}

func init() {
    f := BDharmrajeshwar达摩罗逝湿伐罗.(*达摩罗逝湿伐罗DharmrajeshwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharmrajeshwar",
		TitleName: "达摩罗逝湿伐罗",
		TitleCode: "b_dharmrajeshwar",
	}
}
