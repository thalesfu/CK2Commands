package chartres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷蒂尼BretignyBarony struct {
	feud.BaseBarony
}

var BBretigny布雷蒂尼 feud.Barony = &布雷蒂尼BretignyBarony{}

func init() {
	f := BBretigny布雷蒂尼.(*布雷蒂尼BretignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bretigny",
		TitleName: "布雷蒂尼",
		TitleCode: "b_bretigny",
	}
}
