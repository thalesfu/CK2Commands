package lecce

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥里亚OriaBarony struct {
	feud.BaseBarony
}

var BOria奥里亚 feud.Barony = &奥里亚OriaBarony{}

func init() {
	f := BOria奥里亚.(*奥里亚OriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oria",
		TitleName: "奥里亚",
		TitleCode: "b_oria",
	}
}
