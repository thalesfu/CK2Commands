package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒哈亚Kyzyl_khayaBarony struct {
	feud.BaseBarony
}

var BKyzyl_khaya克孜勒哈亚 feud.Barony = &克孜勒哈亚Kyzyl_khayaBarony{}

func init() {
    f := BKyzyl_khaya克孜勒哈亚.(*克孜勒哈亚Kyzyl_khayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl_khaya",
		TitleName: "克孜勒哈亚",
		TitleCode: "b_kyzyl_khaya",
	}
}
