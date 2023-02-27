package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥克拉德尼科瓦OkladnikowaBarony struct {
	feud.BaseBarony
}

var BOkladnikowa奥克拉德尼科瓦 feud.Barony = &奥克拉德尼科瓦OkladnikowaBarony{}

func init() {
    f := BOkladnikowa奥克拉德尼科瓦.(*奥克拉德尼科瓦OkladnikowaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "okladnikowa",
		TitleName: "奥克拉德尼科瓦",
		TitleCode: "b_okladnikowa",
	}
}
