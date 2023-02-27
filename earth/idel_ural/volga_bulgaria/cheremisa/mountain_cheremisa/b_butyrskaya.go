package mountain_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布特尔斯卡亚ButyrskayaBarony struct {
	feud.BaseBarony
}

var BButyrskaya布特尔斯卡亚 feud.Barony = &布特尔斯卡亚ButyrskayaBarony{}

func init() {
    f := BButyrskaya布特尔斯卡亚.(*布特尔斯卡亚ButyrskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "butyrskaya",
		TitleName: "布特尔斯卡亚",
		TitleCode: "b_butyrskaya",
	}
}
