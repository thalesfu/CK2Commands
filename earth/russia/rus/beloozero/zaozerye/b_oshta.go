package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥什塔OshtaBarony struct {
	feud.BaseBarony
}

var BOshta奥什塔 feud.Barony = &奥什塔OshtaBarony{}

func init() {
	f := BOshta奥什塔.(*奥什塔OshtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oshta",
		TitleName: "奥什塔",
		TitleCode: "b_oshta",
	}
}
