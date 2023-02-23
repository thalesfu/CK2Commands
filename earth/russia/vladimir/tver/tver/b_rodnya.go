package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗德尼亚RodnyaBarony struct {
	feud.BaseBarony
}

var BRodnya罗德尼亚 feud.Barony = &罗德尼亚RodnyaBarony{}

func init() {
	f := BRodnya罗德尼亚.(*罗德尼亚RodnyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rodnya",
		TitleName: "罗德尼亚",
		TitleCode: "b_rodnya",
	}
}
