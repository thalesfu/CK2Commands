package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨比奥内塔SabbionetaBarony struct {
	feud.BaseBarony
}

var BSabbioneta萨比奥内塔 feud.Barony = &萨比奥内塔SabbionetaBarony{}

func init() {
	f := BSabbioneta萨比奥内塔.(*萨比奥内塔SabbionetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sabbioneta",
		TitleName: "萨比奥内塔",
		TitleCode: "b_sabbioneta",
	}
}
