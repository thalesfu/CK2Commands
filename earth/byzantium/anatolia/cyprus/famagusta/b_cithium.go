package famagusta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基蒂翁CithiumBarony struct {
	feud.BaseBarony
}

var BCithium基蒂翁 feud.Barony = &基蒂翁CithiumBarony{}

func init() {
	f := BCithium基蒂翁.(*基蒂翁CithiumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cithium",
		TitleName: "基蒂翁",
		TitleCode: "b_cithium",
	}
}
