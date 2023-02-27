package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黄草坝HwantsaopaBarony struct {
	feud.BaseBarony
}

var BHwantsaopa黄草坝 feud.Barony = &黄草坝HwantsaopaBarony{}

func init() {
    f := BHwantsaopa黄草坝.(*黄草坝HwantsaopaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hwantsaopa",
		TitleName: "黄草坝",
		TitleCode: "b_hwantsaopa",
	}
}
