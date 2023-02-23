package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦利德沃尔VeldvorBarony struct {
	feud.BaseBarony
}

var BVeldvor韦利德沃尔 feud.Barony = &韦利德沃尔VeldvorBarony{}

func init() {
	f := BVeldvor韦利德沃尔.(*韦利德沃尔VeldvorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veldvor",
		TitleName: "韦利德沃尔",
		TitleCode: "b_veldvor",
	}
}
