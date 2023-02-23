package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉迪科法尼RadicofaniBarony struct {
	feud.BaseBarony
}

var BRadicofani拉迪科法尼 feud.Barony = &拉迪科法尼RadicofaniBarony{}

func init() {
	f := BRadicofani拉迪科法尼.(*拉迪科法尼RadicofaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radicofani",
		TitleName: "拉迪科法尼",
		TitleCode: "b_radicofani",
	}
}
