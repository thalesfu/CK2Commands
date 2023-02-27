package talava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维伊齐埃姆斯VijciemsBarony struct {
	feud.BaseBarony
}

var BVijciems维伊齐埃姆斯 feud.Barony = &维伊齐埃姆斯VijciemsBarony{}

func init() {
    f := BVijciems维伊齐埃姆斯.(*维伊齐埃姆斯VijciemsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vijciems",
		TitleName: "维伊齐埃姆斯",
		TitleCode: "b_vijciems",
	}
}
