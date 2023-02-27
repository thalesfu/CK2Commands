package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇什梅ChishmyBarony struct {
	feud.BaseBarony
}

var BChishmy奇什梅 feud.Barony = &奇什梅ChishmyBarony{}

func init() {
    f := BChishmy奇什梅.(*奇什梅ChishmyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chishmy",
		TitleName: "奇什梅",
		TitleCode: "b_chishmy",
	}
}
