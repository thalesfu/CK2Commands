package perigord

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔拉SarlatBarony struct {
	feud.BaseBarony
}

var BSarlat萨尔拉 feud.Barony = &萨尔拉SarlatBarony{}

func init() {
    f := BSarlat萨尔拉.(*萨尔拉SarlatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarlat",
		TitleName: "萨尔拉",
		TitleCode: "b_sarlat",
	}
}
