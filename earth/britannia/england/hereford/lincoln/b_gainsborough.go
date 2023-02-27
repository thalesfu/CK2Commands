package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖恩斯伯勒GainsboroughBarony struct {
	feud.BaseBarony
}

var BGainsborough盖恩斯伯勒 feud.Barony = &盖恩斯伯勒GainsboroughBarony{}

func init() {
    f := BGainsborough盖恩斯伯勒.(*盖恩斯伯勒GainsboroughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gainsborough",
		TitleName: "盖恩斯伯勒",
		TitleCode: "b_gainsborough",
	}
}
