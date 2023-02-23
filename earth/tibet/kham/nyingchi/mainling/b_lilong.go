package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里龙LilongBarony struct {
	feud.BaseBarony
}

var BLilong里龙 feud.Barony = &里龙LilongBarony{}

func init() {
	f := BLilong里龙.(*里龙LilongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lilong",
		TitleName: "里龙",
		TitleCode: "b_lilong",
	}
}
