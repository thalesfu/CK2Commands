package rovaniemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基蒂宁KitinenBarony struct {
	feud.BaseBarony
}

var BKitinen基蒂宁 feud.Barony = &基蒂宁KitinenBarony{}

func init() {
    f := BKitinen基蒂宁.(*基蒂宁KitinenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kitinen",
		TitleName: "基蒂宁",
		TitleCode: "b_kitinen",
	}
}
