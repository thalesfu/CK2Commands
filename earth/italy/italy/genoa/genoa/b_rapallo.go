package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉帕洛RapalloBarony struct {
	feud.BaseBarony
}

var BRapallo拉帕洛 feud.Barony = &拉帕洛RapalloBarony{}

func init() {
    f := BRapallo拉帕洛.(*拉帕洛RapalloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rapallo",
		TitleName: "拉帕洛",
		TitleCode: "b_rapallo",
	}
}
