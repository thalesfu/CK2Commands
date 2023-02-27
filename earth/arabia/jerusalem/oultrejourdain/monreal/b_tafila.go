package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔菲拉TafilaBarony struct {
	feud.BaseBarony
}

var BTafila塔菲拉 feud.Barony = &塔菲拉TafilaBarony{}

func init() {
    f := BTafila塔菲拉.(*塔菲拉TafilaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tafila",
		TitleName: "塔菲拉",
		TitleCode: "b_tafila",
	}
}
