package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛奇科沃KlochkovoBarony struct {
	feud.BaseBarony
}

var BKlochkovo克洛奇科沃 feud.Barony = &克洛奇科沃KlochkovoBarony{}

func init() {
	f := BKlochkovo克洛奇科沃.(*克洛奇科沃KlochkovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klochkovo",
		TitleName: "克洛奇科沃",
		TitleCode: "b_klochkovo",
	}
}
