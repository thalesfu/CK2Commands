package sortavala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马兰KarmalanBarony struct {
	feud.BaseBarony
}

var BKarmalan卡马兰 feud.Barony = &卡马兰KarmalanBarony{}

func init() {
	f := BKarmalan卡马兰.(*卡马兰KarmalanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karmalan",
		TitleName: "卡马兰",
		TitleCode: "b_karmalan",
	}
}
