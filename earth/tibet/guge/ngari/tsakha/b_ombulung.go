package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 屋木龙OmbulungBarony struct {
	feud.BaseBarony
}

var BOmbulung屋木龙 feud.Barony = &屋木龙OmbulungBarony{}

func init() {
	f := BOmbulung屋木龙.(*屋木龙OmbulungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ombulung",
		TitleName: "屋木龙",
		TitleCode: "b_ombulung",
	}
}
