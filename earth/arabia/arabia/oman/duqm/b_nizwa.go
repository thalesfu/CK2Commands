package duqm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈兹瓦NizwaBarony struct {
	feud.BaseBarony
}

var BNizwa奈兹瓦 feud.Barony = &奈兹瓦NizwaBarony{}

func init() {
	f := BNizwa奈兹瓦.(*奈兹瓦NizwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizwa",
		TitleName: "奈兹瓦",
		TitleCode: "b_nizwa",
	}
}
