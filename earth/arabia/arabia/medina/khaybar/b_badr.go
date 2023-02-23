package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜德尔BadrBarony struct {
	feud.BaseBarony
}

var BBadr拜德尔 feud.Barony = &拜德尔BadrBarony{}

func init() {
	f := BBadr拜德尔.(*拜德尔BadrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badr",
		TitleName: "拜德尔",
		TitleCode: "b_badr",
	}
}
