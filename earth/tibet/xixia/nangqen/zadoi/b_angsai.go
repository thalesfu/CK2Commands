package zadoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂赛AngsaiBarony struct {
	feud.BaseBarony
}

var BAngsai昂赛 feud.Barony = &昂赛AngsaiBarony{}

func init() {
    f := BAngsai昂赛.(*昂赛AngsaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "angsai",
		TitleName: "昂赛",
		TitleCode: "b_angsai",
	}
}
