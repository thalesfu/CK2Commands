package thessalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特里KastriBarony struct {
	feud.BaseBarony
}

var BKastri卡斯特里 feud.Barony = &卡斯特里KastriBarony{}

func init() {
	f := BKastri卡斯特里.(*卡斯特里KastriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastri",
		TitleName: "卡斯特里",
		TitleCode: "b_kastri",
	}
}
