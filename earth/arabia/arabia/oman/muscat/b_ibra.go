package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊卜拉IbraBarony struct {
	feud.BaseBarony
}

var BIbra伊卜拉 feud.Barony = &伊卜拉IbraBarony{}

func init() {
	f := BIbra伊卜拉.(*伊卜拉IbraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ibra",
		TitleName: "伊卜拉",
		TitleCode: "b_ibra",
	}
}
