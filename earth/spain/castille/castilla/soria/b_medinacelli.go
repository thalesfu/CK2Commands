package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅迪纳塞利MedinacelliBarony struct {
	feud.BaseBarony
}

var BMedinacelli梅迪纳塞利 feud.Barony = &梅迪纳塞利MedinacelliBarony{}

func init() {
	f := BMedinacelli梅迪纳塞利.(*梅迪纳塞利MedinacelliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medinacelli",
		TitleName: "梅迪纳塞利",
		TitleCode: "b_medinacelli",
	}
}
