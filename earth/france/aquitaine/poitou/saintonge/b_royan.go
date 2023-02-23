package saintonge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁瓦扬RoyanBarony struct {
	feud.BaseBarony
}

var BRoyan鲁瓦扬 feud.Barony = &鲁瓦扬RoyanBarony{}

func init() {
	f := BRoyan鲁瓦扬.(*鲁瓦扬RoyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "royan",
		TitleName: "鲁瓦扬",
		TitleCode: "b_royan",
	}
}
