package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔莱塔BarlettaBarony struct {
	feud.BaseBarony
}

var BBarletta巴尔莱塔 feud.Barony = &巴尔莱塔BarlettaBarony{}

func init() {
    f := BBarletta巴尔莱塔.(*巴尔莱塔BarlettaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barletta",
		TitleName: "巴尔莱塔",
		TitleCode: "b_barletta",
	}
}
