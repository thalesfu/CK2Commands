package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比韦罗ViveiroBarony struct {
	feud.BaseBarony
}

var BViveiro比韦罗 feud.Barony = &比韦罗ViveiroBarony{}

func init() {
    f := BViveiro比韦罗.(*比韦罗ViveiroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viveiro",
		TitleName: "比韦罗",
		TitleCode: "b_viveiro",
	}
}
