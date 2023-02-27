package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别日奇BezichyBarony struct {
	feud.BaseBarony
}

var BBezichy别日奇 feud.Barony = &别日奇BezichyBarony{}

func init() {
    f := BBezichy别日奇.(*别日奇BezichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bezichy",
		TitleName: "别日奇",
		TitleCode: "b_bezichy",
	}
}
