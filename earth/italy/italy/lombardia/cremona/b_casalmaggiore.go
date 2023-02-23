package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡萨尔马焦雷CasalmaggioreBarony struct {
	feud.BaseBarony
}

var BCasalmaggiore卡萨尔马焦雷 feud.Barony = &卡萨尔马焦雷CasalmaggioreBarony{}

func init() {
	f := BCasalmaggiore卡萨尔马焦雷.(*卡萨尔马焦雷CasalmaggioreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "casalmaggiore",
		TitleName: "卡萨尔马焦雷",
		TitleCode: "b_casalmaggiore",
	}
}
