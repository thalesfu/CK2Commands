package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯克雷RoscreaBarony struct {
	feud.BaseBarony
}

var BRoscrea罗斯克雷 feud.Barony = &罗斯克雷RoscreaBarony{}

func init() {
    f := BRoscrea罗斯克雷.(*罗斯克雷RoscreaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roscrea",
		TitleName: "罗斯克雷",
		TitleCode: "b_roscrea",
	}
}
