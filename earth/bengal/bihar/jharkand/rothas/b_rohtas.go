package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢诃多娑RohtasBarony struct {
	feud.BaseBarony
}

var BRohtas卢诃多娑 feud.Barony = &卢诃多娑RohtasBarony{}

func init() {
	f := BRohtas卢诃多娑.(*卢诃多娑RohtasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rohtas",
		TitleName: "卢诃多娑",
		TitleCode: "b_rohtas",
	}
}
