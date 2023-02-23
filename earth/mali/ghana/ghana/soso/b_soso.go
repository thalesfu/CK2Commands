package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索索SosoBarony struct {
	feud.BaseBarony
}

var BSoso索索 feud.Barony = &索索SosoBarony{}

func init() {
	f := BSoso索索.(*索索SosoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soso",
		TitleName: "索索",
		TitleCode: "b_soso",
	}
}
