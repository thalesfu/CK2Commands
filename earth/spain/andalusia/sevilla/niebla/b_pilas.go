package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮拉斯PilasBarony struct {
	feud.BaseBarony
}

var BPilas皮拉斯 feud.Barony = &皮拉斯PilasBarony{}

func init() {
	f := BPilas皮拉斯.(*皮拉斯PilasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pilas",
		TitleName: "皮拉斯",
		TitleCode: "b_pilas",
	}
}
