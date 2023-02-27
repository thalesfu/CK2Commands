package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔顿HaltonBarony struct {
	feud.BaseBarony
}

var BHalton哈尔顿 feud.Barony = &哈尔顿HaltonBarony{}

func init() {
    f := BHalton哈尔顿.(*哈尔顿HaltonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halton",
		TitleName: "哈尔顿",
		TitleCode: "b_halton",
	}
}
