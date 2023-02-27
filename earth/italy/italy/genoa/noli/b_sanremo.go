package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣雷莫SanremoBarony struct {
	feud.BaseBarony
}

var BSanremo圣雷莫 feud.Barony = &圣雷莫SanremoBarony{}

func init() {
    f := BSanremo圣雷莫.(*圣雷莫SanremoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanremo",
		TitleName: "圣雷莫",
		TitleCode: "b_sanremo",
	}
}
