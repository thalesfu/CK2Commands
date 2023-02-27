package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊卜IbbBarony struct {
	feud.BaseBarony
}

var BIbb伊卜 feud.Barony = &伊卜IbbBarony{}

func init() {
    f := BIbb伊卜.(*伊卜IbbBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ibb",
		TitleName: "伊卜",
		TitleCode: "b_ibb",
	}
}
