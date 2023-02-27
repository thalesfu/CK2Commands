package plauen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅泽堡MerseburgBarony struct {
	feud.BaseBarony
}

var BMerseburg梅泽堡 feud.Barony = &梅泽堡MerseburgBarony{}

func init() {
    f := BMerseburg梅泽堡.(*梅泽堡MerseburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "merseburg",
		TitleName: "梅泽堡",
		TitleCode: "b_merseburg",
	}
}
