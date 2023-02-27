package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯罗拏罗KarnalBarony struct {
	feud.BaseBarony
}

var BKarnal羯罗拏罗 feud.Barony = &羯罗拏罗KarnalBarony{}

func init() {
    f := BKarnal羯罗拏罗.(*羯罗拏罗KarnalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karnal",
		TitleName: "羯罗拏罗",
		TitleCode: "b_karnal",
	}
}
