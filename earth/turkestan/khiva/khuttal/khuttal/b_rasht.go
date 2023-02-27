package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉什特RashtBarony struct {
	feud.BaseBarony
}

var BRasht拉什特 feud.Barony = &拉什特RashtBarony{}

func init() {
    f := BRasht拉什特.(*拉什特RashtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rasht",
		TitleName: "拉什特",
		TitleCode: "b_rasht",
	}
}
