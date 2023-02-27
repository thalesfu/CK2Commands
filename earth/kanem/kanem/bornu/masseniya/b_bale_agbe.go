package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴莱阿贝Bale_agbeBarony struct {
	feud.BaseBarony
}

var BBale_agbe巴莱阿贝 feud.Barony = &巴莱阿贝Bale_agbeBarony{}

func init() {
    f := BBale_agbe巴莱阿贝.(*巴莱阿贝Bale_agbeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bale_agbe",
		TitleName: "巴莱阿贝",
		TitleCode: "b_bale_agbe",
	}
}
