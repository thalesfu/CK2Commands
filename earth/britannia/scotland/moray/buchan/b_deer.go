package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪尔DeerBarony struct {
	feud.BaseBarony
}

var BDeer迪尔 feud.Barony = &迪尔DeerBarony{}

func init() {
	f := BDeer迪尔.(*迪尔DeerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deer",
		TitleName: "迪尔",
		TitleCode: "b_deer",
	}
}
