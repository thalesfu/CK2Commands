package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希恩SkienBarony struct {
	feud.BaseBarony
}

var BSkien希恩 feud.Barony = &希恩SkienBarony{}

func init() {
    f := BSkien希恩.(*希恩SkienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skien",
		TitleName: "希恩",
		TitleCode: "b_skien",
	}
}
