package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯纳哈尔IznajarBarony struct {
	feud.BaseBarony
}

var BIznajar伊斯纳哈尔 feud.Barony = &伊斯纳哈尔IznajarBarony{}

func init() {
	f := BIznajar伊斯纳哈尔.(*伊斯纳哈尔IznajarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iznajar",
		TitleName: "伊斯纳哈尔",
		TitleCode: "b_iznajar",
	}
}
