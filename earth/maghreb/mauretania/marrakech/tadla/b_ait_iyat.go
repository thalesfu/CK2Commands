package tadla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾特伊亚特Ait_iyatBarony struct {
	feud.BaseBarony
}

var BAit_iyat艾特伊亚特 feud.Barony = &艾特伊亚特Ait_iyatBarony{}

func init() {
    f := BAit_iyat艾特伊亚特.(*艾特伊亚特Ait_iyatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ait_iyat",
		TitleName: "艾特伊亚特",
		TitleCode: "b_ait_iyat",
	}
}
