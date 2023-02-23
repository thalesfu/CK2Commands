package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛斯特新堡KlosterneuburgBarony struct {
	feud.BaseBarony
}

var BKlosterneuburg克洛斯特新堡 feud.Barony = &克洛斯特新堡KlosterneuburgBarony{}

func init() {
	f := BKlosterneuburg克洛斯特新堡.(*克洛斯特新堡KlosterneuburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klosterneuburg",
		TitleName: "克洛斯特新堡",
		TitleCode: "b_klosterneuburg",
	}
}
