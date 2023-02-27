package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富尔纳里FurnariBarony struct {
	feud.BaseBarony
}

var BFurnari富尔纳里 feud.Barony = &富尔纳里FurnariBarony{}

func init() {
    f := BFurnari富尔纳里.(*富尔纳里FurnariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furnari",
		TitleName: "富尔纳里",
		TitleCode: "b_furnari",
	}
}
