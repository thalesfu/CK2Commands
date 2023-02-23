package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里辛RuthinBarony struct {
	feud.BaseBarony
}

var BRuthin里辛 feud.Barony = &里辛RuthinBarony{}

func init() {
	f := BRuthin里辛.(*里辛RuthinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruthin",
		TitleName: "里辛",
		TitleCode: "b_ruthin",
	}
}
