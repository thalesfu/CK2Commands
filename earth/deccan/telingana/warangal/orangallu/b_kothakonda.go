package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘多军荼KothakondaBarony struct {
	feud.BaseBarony
}

var BKothakonda拘多军荼 feud.Barony = &拘多军荼KothakondaBarony{}

func init() {
    f := BKothakonda拘多军荼.(*拘多军荼KothakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kothakonda",
		TitleName: "拘多军荼",
		TitleCode: "b_kothakonda",
	}
}
