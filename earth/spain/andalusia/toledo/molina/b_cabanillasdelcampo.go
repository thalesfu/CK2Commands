package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡瓦尼利亚斯德尔坎波CabanillasdelcampoBarony struct {
	feud.BaseBarony
}

var BCabanillasdelcampo卡瓦尼利亚斯德尔坎波 feud.Barony = &卡瓦尼利亚斯德尔坎波CabanillasdelcampoBarony{}

func init() {
	f := BCabanillasdelcampo卡瓦尼利亚斯德尔坎波.(*卡瓦尼利亚斯德尔坎波CabanillasdelcampoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cabanillasdelcampo",
		TitleName: "卡瓦尼利亚斯德尔坎波",
		TitleCode: "b_cabanillasdelcampo",
	}
}
