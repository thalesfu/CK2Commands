package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里亚斯诺耶RyasnoyeBarony struct {
	feud.BaseBarony
}

var BRyasnoye里亚斯诺耶 feud.Barony = &里亚斯诺耶RyasnoyeBarony{}

func init() {
    f := BRyasnoye里亚斯诺耶.(*里亚斯诺耶RyasnoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ryasnoye",
		TitleName: "里亚斯诺耶",
		TitleCode: "b_ryasnoye",
	}
}
