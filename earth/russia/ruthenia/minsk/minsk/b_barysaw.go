package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍里索夫BarysawBarony struct {
	feud.BaseBarony
}

var BBarysaw鲍里索夫 feud.Barony = &鲍里索夫BarysawBarony{}

func init() {
    f := BBarysaw鲍里索夫.(*鲍里索夫BarysawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barysaw",
		TitleName: "鲍里索夫",
		TitleCode: "b_barysaw",
	}
}
