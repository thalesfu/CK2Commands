package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔格萨HargeysaBarony struct {
	feud.BaseBarony
}

var BHargeysa哈尔格萨 feud.Barony = &哈尔格萨HargeysaBarony{}

func init() {
	f := BHargeysa哈尔格萨.(*哈尔格萨HargeysaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hargeysa",
		TitleName: "哈尔格萨",
		TitleCode: "b_hargeysa",
	}
}
