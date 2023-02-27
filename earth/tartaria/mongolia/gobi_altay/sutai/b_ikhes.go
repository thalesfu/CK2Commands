package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊赫斯IkhesBarony struct {
	feud.BaseBarony
}

var BIkhes伊赫斯 feud.Barony = &伊赫斯IkhesBarony{}

func init() {
    f := BIkhes伊赫斯.(*伊赫斯IkhesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikhes",
		TitleName: "伊赫斯",
		TitleCode: "b_ikhes",
	}
}
