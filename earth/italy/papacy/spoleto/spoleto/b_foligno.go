package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福利尼奥FolignoBarony struct {
	feud.BaseBarony
}

var BFoligno福利尼奥 feud.Barony = &福利尼奥FolignoBarony{}

func init() {
	f := BFoligno福利尼奥.(*福利尼奥FolignoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "foligno",
		TitleName: "福利尼奥",
		TitleCode: "b_foligno",
	}
}
