package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列尼Reni_roBarony struct {
	feud.BaseBarony
}

var BReni_ro列尼 feud.Barony = &列尼Reni_roBarony{}

func init() {
    f := BReni_ro列尼.(*列尼Reni_roBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reni_ro",
		TitleName: "列尼",
		TitleCode: "b_reni_ro",
	}
}
