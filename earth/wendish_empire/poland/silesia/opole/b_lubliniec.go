package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢布利涅茨LubliniecBarony struct {
	feud.BaseBarony
}

var BLubliniec卢布利涅茨 feud.Barony = &卢布利涅茨LubliniecBarony{}

func init() {
    f := BLubliniec卢布利涅茨.(*卢布利涅茨LubliniecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lubliniec",
		TitleName: "卢布利涅茨",
		TitleCode: "b_lubliniec",
	}
}
