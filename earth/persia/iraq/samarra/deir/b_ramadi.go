package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉马迪RamadiBarony struct {
	feud.BaseBarony
}

var BRamadi拉马迪 feud.Barony = &拉马迪RamadiBarony{}

func init() {
    f := BRamadi拉马迪.(*拉马迪RamadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramadi",
		TitleName: "拉马迪",
		TitleCode: "b_ramadi",
	}
}
