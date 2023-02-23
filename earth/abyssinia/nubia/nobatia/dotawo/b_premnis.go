package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普雷姆尼斯PremnisBarony struct {
	feud.BaseBarony
}

var BPremnis普雷姆尼斯 feud.Barony = &普雷姆尼斯PremnisBarony{}

func init() {
	f := BPremnis普雷姆尼斯.(*普雷姆尼斯PremnisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "premnis",
		TitleName: "普雷姆尼斯",
		TitleCode: "b_premnis",
	}
}
