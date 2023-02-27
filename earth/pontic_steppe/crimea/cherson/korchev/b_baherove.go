package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴格罗韦BaheroveBarony struct {
	feud.BaseBarony
}

var BBaherove巴格罗韦 feud.Barony = &巴格罗韦BaheroveBarony{}

func init() {
    f := BBaherove巴格罗韦.(*巴格罗韦BaheroveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baherove",
		TitleName: "巴格罗韦",
		TitleCode: "b_baherove",
	}
}
