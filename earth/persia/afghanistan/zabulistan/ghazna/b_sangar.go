package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑加尔SangarBarony struct {
	feud.BaseBarony
}

var BSangar桑加尔 feud.Barony = &桑加尔SangarBarony{}

func init() {
	f := BSangar桑加尔.(*桑加尔SangarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangar",
		TitleName: "桑加尔",
		TitleCode: "b_sangar",
	}
}
