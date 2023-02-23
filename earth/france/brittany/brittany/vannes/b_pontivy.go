package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬蒂维PontivyBarony struct {
	feud.BaseBarony
}

var BPontivy蓬蒂维 feud.Barony = &蓬蒂维PontivyBarony{}

func init() {
	f := BPontivy蓬蒂维.(*蓬蒂维PontivyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pontivy",
		TitleName: "蓬蒂维",
		TitleCode: "b_pontivy",
	}
}
