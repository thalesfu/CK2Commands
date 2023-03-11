package sieradzko-leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 华沙WarszawaBarony struct {
	feud.BaseBarony
}

var BWarszawa华沙 feud.Barony = &华沙WarszawaBarony{}

func init() {
    f := BWarszawa华沙.(*华沙WarszawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "warszawa",
		TitleName: "华沙",
		TitleCode: "b_warszawa",
	}
}
