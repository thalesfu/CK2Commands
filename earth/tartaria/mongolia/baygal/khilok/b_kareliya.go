package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡列利亚KareliyaBarony struct {
	feud.BaseBarony
}

var BKareliya卡列利亚 feud.Barony = &卡列利亚KareliyaBarony{}

func init() {
	f := BKareliya卡列利亚.(*卡列利亚KareliyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kareliya",
		TitleName: "卡列利亚",
		TitleCode: "b_kareliya",
	}
}
