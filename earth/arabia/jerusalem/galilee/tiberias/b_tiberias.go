package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 太巴列TiberiasBarony struct {
	feud.BaseBarony
}

var BTiberias太巴列 feud.Barony = &太巴列TiberiasBarony{}

func init() {
	f := BTiberias太巴列.(*太巴列TiberiasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiberias",
		TitleName: "太巴列",
		TitleCode: "b_tiberias",
	}
}
