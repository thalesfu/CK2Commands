package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦利戈斯蒂VeligostiBarony struct {
	feud.BaseBarony
}

var BVeligosti韦利戈斯蒂 feud.Barony = &韦利戈斯蒂VeligostiBarony{}

func init() {
    f := BVeligosti韦利戈斯蒂.(*韦利戈斯蒂VeligostiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veligosti",
		TitleName: "韦利戈斯蒂",
		TitleCode: "b_veligosti",
	}
}
