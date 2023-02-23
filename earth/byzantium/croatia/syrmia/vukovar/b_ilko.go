package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔科IlkoBarony struct {
	feud.BaseBarony
}

var BIlko伊尔科 feud.Barony = &伊尔科IlkoBarony{}

func init() {
	f := BIlko伊尔科.(*伊尔科IlkoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilko",
		TitleName: "伊尔科",
		TitleCode: "b_ilko",
	}
}
