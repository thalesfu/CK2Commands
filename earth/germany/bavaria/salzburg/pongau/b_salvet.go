package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔韦特SalvetBarony struct {
	feud.BaseBarony
}

var BSalvet萨尔韦特 feud.Barony = &萨尔韦特SalvetBarony{}

func init() {
    f := BSalvet萨尔韦特.(*萨尔韦特SalvetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salvet",
		TitleName: "萨尔韦特",
		TitleCode: "b_salvet",
	}
}
