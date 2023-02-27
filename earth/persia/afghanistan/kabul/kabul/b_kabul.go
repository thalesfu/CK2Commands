package kabul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦布罗KabulBarony struct {
	feud.BaseBarony
}

var BKabul迦布罗 feud.Barony = &迦布罗KabulBarony{}

func init() {
    f := BKabul迦布罗.(*迦布罗KabulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabul",
		TitleName: "迦布罗",
		TitleCode: "b_kabul",
	}
}
