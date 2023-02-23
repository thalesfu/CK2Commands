package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 仙石XianshiBarony struct {
	feud.BaseBarony
}

var BXianshi仙石 feud.Barony = &仙石XianshiBarony{}

func init() {
	f := BXianshi仙石.(*仙石XianshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xianshi",
		TitleName: "仙石",
		TitleCode: "b_xianshi",
	}
}
