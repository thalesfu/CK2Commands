package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔津BurqinBarony struct {
	feud.BaseBarony
}

var BBurqin布尔津 feud.Barony = &布尔津BurqinBarony{}

func init() {
    f := BBurqin布尔津.(*布尔津BurqinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burqin",
		TitleName: "布尔津",
		TitleCode: "b_burqin",
	}
}
