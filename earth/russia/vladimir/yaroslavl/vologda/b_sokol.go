package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索科尔SokolBarony struct {
	feud.BaseBarony
}

var BSokol索科尔 feud.Barony = &索科尔SokolBarony{}

func init() {
	f := BSokol索科尔.(*索科尔SokolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sokol",
		TitleName: "索科尔",
		TitleCode: "b_sokol",
	}
}
