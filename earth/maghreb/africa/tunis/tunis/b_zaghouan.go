package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宰格万ZaghouanBarony struct {
	feud.BaseBarony
}

var BZaghouan宰格万 feud.Barony = &宰格万ZaghouanBarony{}

func init() {
	f := BZaghouan宰格万.(*宰格万ZaghouanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaghouan",
		TitleName: "宰格万",
		TitleCode: "b_zaghouan",
	}
}
