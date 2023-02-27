package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼拜杰An_nibajBarony struct {
	feud.BaseBarony
}

var BAn_nibaj尼拜杰 feud.Barony = &尼拜杰An_nibajBarony{}

func init() {
    f := BAn_nibaj尼拜杰.(*尼拜杰An_nibajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "an_nibaj",
		TitleName: "尼拜杰",
		TitleCode: "b_an_nibaj",
	}
}
