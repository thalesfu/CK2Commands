package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯利布林QasribrimBarony struct {
	feud.BaseBarony
}

var BQasribrim卡斯利布林 feud.Barony = &卡斯利布林QasribrimBarony{}

func init() {
	f := BQasribrim卡斯利布林.(*卡斯利布林QasribrimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qasribrim",
		TitleName: "卡斯利布林",
		TitleCode: "b_qasribrim",
	}
}
