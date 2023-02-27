package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔斯卡亚KurskayaBarony struct {
	feud.BaseBarony
}

var BKurskaya库尔斯卡亚 feud.Barony = &库尔斯卡亚KurskayaBarony{}

func init() {
    f := BKurskaya库尔斯卡亚.(*库尔斯卡亚KurskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurskaya",
		TitleName: "库尔斯卡亚",
		TitleCode: "b_kurskaya",
	}
}
