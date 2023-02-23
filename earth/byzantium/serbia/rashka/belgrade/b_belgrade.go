package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔格莱德BelgradeBarony struct {
	feud.BaseBarony
}

var BBelgrade贝尔格莱德 feud.Barony = &贝尔格莱德BelgradeBarony{}

func init() {
	f := BBelgrade贝尔格莱德.(*贝尔格莱德BelgradeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belgrade",
		TitleName: "贝尔格莱德",
		TitleCode: "b_belgrade",
	}
}
