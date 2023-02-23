package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯基宁KaskinenBarony struct {
	feud.BaseBarony
}

var BKaskinen卡斯基宁 feud.Barony = &卡斯基宁KaskinenBarony{}

func init() {
	f := BKaskinen卡斯基宁.(*卡斯基宁KaskinenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaskinen",
		TitleName: "卡斯基宁",
		TitleCode: "b_kaskinen",
	}
}
