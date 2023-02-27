package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦基尔巴扎尔WekilbazarBarony struct {
	feud.BaseBarony
}

var BWekilbazar韦基尔巴扎尔 feud.Barony = &韦基尔巴扎尔WekilbazarBarony{}

func init() {
    f := BWekilbazar韦基尔巴扎尔.(*韦基尔巴扎尔WekilbazarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wekilbazar",
		TitleName: "韦基尔巴扎尔",
		TitleCode: "b_wekilbazar",
	}
}
