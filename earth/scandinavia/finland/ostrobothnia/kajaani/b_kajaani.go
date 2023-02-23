package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡亚尼KajaaniBarony struct {
	feud.BaseBarony
}

var BKajaani卡亚尼 feud.Barony = &卡亚尼KajaaniBarony{}

func init() {
	f := BKajaani卡亚尼.(*卡亚尼KajaaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kajaani",
		TitleName: "卡亚尼",
		TitleCode: "b_kajaani",
	}
}
