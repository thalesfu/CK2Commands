package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗姆尼RomneyBarony struct {
	feud.BaseBarony
}

var BRomney罗姆尼 feud.Barony = &罗姆尼RomneyBarony{}

func init() {
	f := BRomney罗姆尼.(*罗姆尼RomneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romney",
		TitleName: "罗姆尼",
		TitleCode: "b_romney",
	}
}
