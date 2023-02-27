package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊兹IzBarony struct {
	feud.BaseBarony
}

var BIz伊兹 feud.Barony = &伊兹IzBarony{}

func init() {
    f := BIz伊兹.(*伊兹IzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iz",
		TitleName: "伊兹",
		TitleCode: "b_iz",
	}
}
