package oxford

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牛津OxfordBarony struct {
	feud.BaseBarony
}

var BOxford牛津 feud.Barony = &牛津OxfordBarony{}

func init() {
    f := BOxford牛津.(*牛津OxfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oxford",
		TitleName: "牛津",
		TitleCode: "b_oxford",
	}
}
