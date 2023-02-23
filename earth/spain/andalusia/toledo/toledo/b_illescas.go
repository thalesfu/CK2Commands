package toledo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊列斯卡斯IllescasBarony struct {
	feud.BaseBarony
}

var BIllescas伊列斯卡斯 feud.Barony = &伊列斯卡斯IllescasBarony{}

func init() {
	f := BIllescas伊列斯卡斯.(*伊列斯卡斯IllescasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illescas",
		TitleName: "伊列斯卡斯",
		TitleCode: "b_illescas",
	}
}
