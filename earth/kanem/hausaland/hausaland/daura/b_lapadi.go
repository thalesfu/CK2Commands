package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉帕迪LapadiBarony struct {
	feud.BaseBarony
}

var BLapadi拉帕迪 feud.Barony = &拉帕迪LapadiBarony{}

func init() {
	f := BLapadi拉帕迪.(*拉帕迪LapadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lapadi",
		TitleName: "拉帕迪",
		TitleCode: "b_lapadi",
	}
}
