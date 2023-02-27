package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣让巴蒂斯特StjeanbaptisteBarony struct {
	feud.BaseBarony
}

var BStjeanbaptiste圣让巴蒂斯特 feud.Barony = &圣让巴蒂斯特StjeanbaptisteBarony{}

func init() {
    f := BStjeanbaptiste圣让巴蒂斯特.(*圣让巴蒂斯特StjeanbaptisteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stjeanbaptiste",
		TitleName: "圣让巴蒂斯特",
		TitleCode: "b_stjeanbaptiste",
	}
}
