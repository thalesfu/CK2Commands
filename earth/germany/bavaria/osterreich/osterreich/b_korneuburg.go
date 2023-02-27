package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔新堡KorneuburgBarony struct {
	feud.BaseBarony
}

var BKorneuburg科尔新堡 feud.Barony = &科尔新堡KorneuburgBarony{}

func init() {
    f := BKorneuburg科尔新堡.(*科尔新堡KorneuburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korneuburg",
		TitleName: "科尔新堡",
		TitleCode: "b_korneuburg",
	}
}
