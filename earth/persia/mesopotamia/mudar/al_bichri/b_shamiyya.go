package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙米亚ShamiyyaBarony struct {
	feud.BaseBarony
}

var BShamiyya沙米亚 feud.Barony = &沙米亚ShamiyyaBarony{}

func init() {
    f := BShamiyya沙米亚.(*沙米亚ShamiyyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shamiyya",
		TitleName: "沙米亚",
		TitleCode: "b_shamiyya",
	}
}
