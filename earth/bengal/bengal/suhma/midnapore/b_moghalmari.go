package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牟迦罗摩利MoghalmariBarony struct {
	feud.BaseBarony
}

var BMoghalmari牟迦罗摩利 feud.Barony = &牟迦罗摩利MoghalmariBarony{}

func init() {
    f := BMoghalmari牟迦罗摩利.(*牟迦罗摩利MoghalmariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moghalmari",
		TitleName: "牟迦罗摩利",
		TitleCode: "b_moghalmari",
	}
}
