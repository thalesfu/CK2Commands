package kuopio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科塔萨里KotasaariBarony struct {
	feud.BaseBarony
}

var BKotasaari科塔萨里 feud.Barony = &科塔萨里KotasaariBarony{}

func init() {
    f := BKotasaari科塔萨里.(*科塔萨里KotasaariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kotasaari",
		TitleName: "科塔萨里",
		TitleCode: "b_kotasaari",
	}
}
