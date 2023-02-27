package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅韦MeweBarony struct {
	feud.BaseBarony
}

var BMewe梅韦 feud.Barony = &梅韦MeweBarony{}

func init() {
    f := BMewe梅韦.(*梅韦MeweBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mewe",
		TitleName: "梅韦",
		TitleCode: "b_mewe",
	}
}
