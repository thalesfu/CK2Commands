package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下德布尼克DolnidabknikBarony struct {
	feud.BaseBarony
}

var BDolnidabknik下德布尼克 feud.Barony = &下德布尼克DolnidabknikBarony{}

func init() {
	f := BDolnidabknik下德布尼克.(*下德布尼克DolnidabknikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dolnidabknik",
		TitleName: "下德布尼克",
		TitleCode: "b_dolnidabknik",
	}
}
