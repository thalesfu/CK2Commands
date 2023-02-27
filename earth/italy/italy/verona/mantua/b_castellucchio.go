package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰卢基奥CastellucchioBarony struct {
	feud.BaseBarony
}

var BCastellucchio卡斯泰卢基奥 feud.Barony = &卡斯泰卢基奥CastellucchioBarony{}

func init() {
    f := BCastellucchio卡斯泰卢基奥.(*卡斯泰卢基奥CastellucchioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castellucchio",
		TitleName: "卡斯泰卢基奥",
		TitleCode: "b_castellucchio",
	}
}
