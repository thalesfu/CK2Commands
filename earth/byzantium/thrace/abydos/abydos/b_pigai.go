package abydos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮盖伊PigaiBarony struct {
	feud.BaseBarony
}

var BPigai皮盖伊 feud.Barony = &皮盖伊PigaiBarony{}

func init() {
	f := BPigai皮盖伊.(*皮盖伊PigaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pigai",
		TitleName: "皮盖伊",
		TitleCode: "b_pigai",
	}
}
