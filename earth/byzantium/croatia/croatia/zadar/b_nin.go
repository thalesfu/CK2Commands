package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼恩NinBarony struct {
	feud.BaseBarony
}

var BNin尼恩 feud.Barony = &尼恩NinBarony{}

func init() {
	f := BNin尼恩.(*尼恩NinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nin",
		TitleName: "尼恩",
		TitleCode: "b_nin",
	}
}
