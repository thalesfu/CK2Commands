package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆悉厥BarkulBarony struct {
	feud.BaseBarony
}

var BBarkul婆悉厥 feud.Barony = &婆悉厥BarkulBarony{}

func init() {
    f := BBarkul婆悉厥.(*婆悉厥BarkulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barkul",
		TitleName: "婆悉厥",
		TitleCode: "b_barkul",
	}
}
