package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔托斯MartosBarony struct {
	feud.BaseBarony
}

var BMartos马尔托斯 feud.Barony = &马尔托斯MartosBarony{}

func init() {
    f := BMartos马尔托斯.(*马尔托斯MartosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "martos",
		TitleName: "马尔托斯",
		TitleCode: "b_martos",
	}
}
