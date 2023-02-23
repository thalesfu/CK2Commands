package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨维利亚诺SaviglianoBarony struct {
	feud.BaseBarony
}

var BSavigliano萨维利亚诺 feud.Barony = &萨维利亚诺SaviglianoBarony{}

func init() {
	f := BSavigliano萨维利亚诺.(*萨维利亚诺SaviglianoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "savigliano",
		TitleName: "萨维利亚诺",
		TitleCode: "b_savigliano",
	}
}
