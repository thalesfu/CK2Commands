package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞拉韦扎SeravezzaBarony struct {
	feud.BaseBarony
}

var BSeravezza塞拉韦扎 feud.Barony = &塞拉韦扎SeravezzaBarony{}

func init() {
    f := BSeravezza塞拉韦扎.(*塞拉韦扎SeravezzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seravezza",
		TitleName: "塞拉韦扎",
		TitleCode: "b_seravezza",
	}
}
