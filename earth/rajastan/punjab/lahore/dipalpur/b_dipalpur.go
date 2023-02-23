package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提波罗补罗DipalpurBarony struct {
	feud.BaseBarony
}

var BDipalpur提波罗补罗 feud.Barony = &提波罗补罗DipalpurBarony{}

func init() {
	f := BDipalpur提波罗补罗.(*提波罗补罗DipalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dipalpur",
		TitleName: "提波罗补罗",
		TitleCode: "b_dipalpur",
	}
}
