package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维泰博ViterboBarony struct {
	feud.BaseBarony
}

var BViterbo维泰博 feud.Barony = &维泰博ViterboBarony{}

func init() {
    f := BViterbo维泰博.(*维泰博ViterboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viterbo",
		TitleName: "维泰博",
		TitleCode: "b_viterbo",
	}
}
