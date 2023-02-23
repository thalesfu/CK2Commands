package taranto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特里卡尼科TricanicoBarony struct {
	feud.BaseBarony
}

var BTricanico特里卡尼科 feud.Barony = &特里卡尼科TricanicoBarony{}

func init() {
	f := BTricanico特里卡尼科.(*特里卡尼科TricanicoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tricanico",
		TitleName: "特里卡尼科",
		TitleCode: "b_tricanico",
	}
}
