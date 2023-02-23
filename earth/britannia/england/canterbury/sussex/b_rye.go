package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉伊RyeBarony struct {
	feud.BaseBarony
}

var BRye拉伊 feud.Barony = &拉伊RyeBarony{}

func init() {
	f := BRye拉伊.(*拉伊RyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rye",
		TitleName: "拉伊",
		TitleCode: "b_rye",
	}
}
