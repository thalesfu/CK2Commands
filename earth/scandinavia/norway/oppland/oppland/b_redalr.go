package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷达尔RedalrBarony struct {
	feud.BaseBarony
}

var BRedalr雷达尔 feud.Barony = &雷达尔RedalrBarony{}

func init() {
    f := BRedalr雷达尔.(*雷达尔RedalrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "redalr",
		TitleName: "雷达尔",
		TitleCode: "b_redalr",
	}
}
