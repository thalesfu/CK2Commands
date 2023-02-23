package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比利亚雷亚尔VilarrealBarony struct {
	feud.BaseBarony
}

var BVilarreal比利亚雷亚尔 feud.Barony = &比利亚雷亚尔VilarrealBarony{}

func init() {
	f := BVilarreal比利亚雷亚尔.(*比利亚雷亚尔VilarrealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vilarreal",
		TitleName: "比利亚雷亚尔",
		TitleCode: "b_vilarreal",
	}
}
