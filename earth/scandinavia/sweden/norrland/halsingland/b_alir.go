package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿利尔AlirBarony struct {
	feud.BaseBarony
}

var BAlir阿利尔 feud.Barony = &阿利尔AlirBarony{}

func init() {
	f := BAlir阿利尔.(*阿利尔AlirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alir",
		TitleName: "阿利尔",
		TitleCode: "b_alir",
	}
}
