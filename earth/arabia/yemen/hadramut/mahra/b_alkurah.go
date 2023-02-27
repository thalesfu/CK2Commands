package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库拉AlkurahBarony struct {
	feud.BaseBarony
}

var BAlkurah库拉 feud.Barony = &库拉AlkurahBarony{}

func init() {
    f := BAlkurah库拉.(*库拉AlkurahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkurah",
		TitleName: "库拉",
		TitleCode: "b_alkurah",
	}
}
