package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊兹拉IzraBarony struct {
	feud.BaseBarony
}

var BIzra伊兹拉 feud.Barony = &伊兹拉IzraBarony{}

func init() {
    f := BIzra伊兹拉.(*伊兹拉IzraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "izra",
		TitleName: "伊兹拉",
		TitleCode: "b_izra",
	}
}
