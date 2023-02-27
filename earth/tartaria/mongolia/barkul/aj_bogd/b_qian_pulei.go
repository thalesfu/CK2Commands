package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 前蒲类Qian_puleiBarony struct {
	feud.BaseBarony
}

var BQian_pulei前蒲类 feud.Barony = &前蒲类Qian_puleiBarony{}

func init() {
    f := BQian_pulei前蒲类.(*前蒲类Qian_puleiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qian_pulei",
		TitleName: "前蒲类",
		TitleCode: "b_qian_pulei",
	}
}
