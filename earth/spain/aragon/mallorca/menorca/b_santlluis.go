package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣路易斯SantlluisBarony struct {
	feud.BaseBarony
}

var BSantlluis圣路易斯 feud.Barony = &圣路易斯SantlluisBarony{}

func init() {
	f := BSantlluis圣路易斯.(*圣路易斯SantlluisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santlluis",
		TitleName: "圣路易斯",
		TitleCode: "b_santlluis",
	}
}
