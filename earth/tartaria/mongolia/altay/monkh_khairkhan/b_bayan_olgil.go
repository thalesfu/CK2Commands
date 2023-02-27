package monkh_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴彦乌列盖Bayan_olgilBarony struct {
	feud.BaseBarony
}

var BBayan_olgil巴彦乌列盖 feud.Barony = &巴彦乌列盖Bayan_olgilBarony{}

func init() {
    f := BBayan_olgil巴彦乌列盖.(*巴彦乌列盖Bayan_olgilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayan_olgil",
		TitleName: "巴彦乌列盖",
		TitleCode: "b_bayan_olgil",
	}
}
