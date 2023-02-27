package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈纳岛NainativuBarony struct {
	feud.BaseBarony
}

var BNainativu奈纳岛 feud.Barony = &奈纳岛NainativuBarony{}

func init() {
    f := BNainativu奈纳岛.(*奈纳岛NainativuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nainativu",
		TitleName: "奈纳岛",
		TitleCode: "b_nainativu",
	}
}
