package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里布尼察RibnicaBarony struct {
	feud.BaseBarony
}

var BRibnica里布尼察 feud.Barony = &里布尼察RibnicaBarony{}

func init() {
	f := BRibnica里布尼察.(*里布尼察RibnicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ribnica",
		TitleName: "里布尼察",
		TitleCode: "b_ribnica",
	}
}
