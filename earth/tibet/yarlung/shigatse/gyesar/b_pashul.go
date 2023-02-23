package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕苏尔PashulBarony struct {
	feud.BaseBarony
}

var BPashul帕苏尔 feud.Barony = &帕苏尔PashulBarony{}

func init() {
	f := BPashul帕苏尔.(*帕苏尔PashulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pashul",
		TitleName: "帕苏尔",
		TitleCode: "b_pashul",
	}
}
