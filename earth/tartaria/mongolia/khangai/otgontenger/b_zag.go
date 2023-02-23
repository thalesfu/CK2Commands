package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎格ZagBarony struct {
	feud.BaseBarony
}

var BZag扎格 feud.Barony = &扎格ZagBarony{}

func init() {
	f := BZag扎格.(*扎格ZagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zag",
		TitleName: "扎格",
		TitleCode: "b_zag",
	}
}
