package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多瑙DunaiBarony struct {
	feud.BaseBarony
}

var BDunai多瑙 feud.Barony = &多瑙DunaiBarony{}

func init() {
    f := BDunai多瑙.(*多瑙DunaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunai",
		TitleName: "多瑙",
		TitleCode: "b_dunai",
	}
}
