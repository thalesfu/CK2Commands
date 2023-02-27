package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 后蒲类Hou_puleiBarony struct {
	feud.BaseBarony
}

var BHou_pulei后蒲类 feud.Barony = &后蒲类Hou_puleiBarony{}

func init() {
    f := BHou_pulei后蒲类.(*后蒲类Hou_puleiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hou_pulei",
		TitleName: "后蒲类",
		TitleCode: "b_hou_pulei",
	}
}
