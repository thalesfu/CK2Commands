package laukaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉卡萨尔米HankasalmiBarony struct {
	feud.BaseBarony
}

var BHankasalmi汉卡萨尔米 feud.Barony = &汉卡萨尔米HankasalmiBarony{}

func init() {
	f := BHankasalmi汉卡萨尔米.(*汉卡萨尔米HankasalmiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hankasalmi",
		TitleName: "汉卡萨尔米",
		TitleCode: "b_hankasalmi",
	}
}
