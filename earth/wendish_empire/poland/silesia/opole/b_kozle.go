package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科伊莱KozleBarony struct {
	feud.BaseBarony
}

var BKozle科伊莱 feud.Barony = &科伊莱KozleBarony{}

func init() {
    f := BKozle科伊莱.(*科伊莱KozleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kozle",
		TitleName: "科伊莱",
		TitleCode: "b_kozle",
	}
}
