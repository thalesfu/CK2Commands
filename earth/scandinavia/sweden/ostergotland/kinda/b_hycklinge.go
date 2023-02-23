package kinda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 许克灵厄HycklingeBarony struct {
	feud.BaseBarony
}

var BHycklinge许克灵厄 feud.Barony = &许克灵厄HycklingeBarony{}

func init() {
	f := BHycklinge许克灵厄.(*许克灵厄HycklingeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hycklinge",
		TitleName: "许克灵厄",
		TitleCode: "b_hycklinge",
	}
}
