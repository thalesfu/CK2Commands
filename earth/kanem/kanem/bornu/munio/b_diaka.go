package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚卡DiakaBarony struct {
	feud.BaseBarony
}

var BDiaka迪亚卡 feud.Barony = &迪亚卡DiakaBarony{}

func init() {
    f := BDiaka迪亚卡.(*迪亚卡DiakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diaka",
		TitleName: "迪亚卡",
		TitleCode: "b_diaka",
	}
}
