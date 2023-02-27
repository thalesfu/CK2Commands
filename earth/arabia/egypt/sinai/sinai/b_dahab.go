package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰海卜DahabBarony struct {
	feud.BaseBarony
}

var BDahab宰海卜 feud.Barony = &宰海卜DahabBarony{}

func init() {
    f := BDahab宰海卜.(*宰海卜DahabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dahab",
		TitleName: "宰海卜",
		TitleCode: "b_dahab",
	}
}
