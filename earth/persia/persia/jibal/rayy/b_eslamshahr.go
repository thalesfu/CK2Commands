package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯兰夏尔EslamshahrBarony struct {
	feud.BaseBarony
}

var BEslamshahr埃斯兰夏尔 feud.Barony = &埃斯兰夏尔EslamshahrBarony{}

func init() {
	f := BEslamshahr埃斯兰夏尔.(*埃斯兰夏尔EslamshahrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eslamshahr",
		TitleName: "埃斯兰夏尔",
		TitleCode: "b_eslamshahr",
	}
}
