package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库兰格KoohrangBarony struct {
	feud.BaseBarony
}

var BKoohrang库兰格 feud.Barony = &库兰格KoohrangBarony{}

func init() {
	f := BKoohrang库兰格.(*库兰格KoohrangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koohrang",
		TitleName: "库兰格",
		TitleCode: "b_koohrang",
	}
}
