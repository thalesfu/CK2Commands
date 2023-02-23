package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉达PradaBarony struct {
	feud.BaseBarony
}

var BPrada普拉达 feud.Barony = &普拉达PradaBarony{}

func init() {
	f := BPrada普拉达.(*普拉达PradaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prada",
		TitleName: "普拉达",
		TitleCode: "b_prada",
	}
}
