package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达纽尔DunureBarony struct {
	feud.BaseBarony
}

var BDunure达纽尔 feud.Barony = &达纽尔DunureBarony{}

func init() {
    f := BDunure达纽尔.(*达纽尔DunureBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunure",
		TitleName: "达纽尔",
		TitleCode: "b_dunure",
	}
}
