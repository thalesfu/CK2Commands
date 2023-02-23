package toropets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇斯托耶ChistoyeBarony struct {
	feud.BaseBarony
}

var BChistoye奇斯托耶 feud.Barony = &奇斯托耶ChistoyeBarony{}

func init() {
	f := BChistoye奇斯托耶.(*奇斯托耶ChistoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chistoye",
		TitleName: "奇斯托耶",
		TitleCode: "b_chistoye",
	}
}
