package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏克阿赫拉斯SoukahrasBarony struct {
	feud.BaseBarony
}

var BSoukahras苏克阿赫拉斯 feud.Barony = &苏克阿赫拉斯SoukahrasBarony{}

func init() {
    f := BSoukahras苏克阿赫拉斯.(*苏克阿赫拉斯SoukahrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soukahras",
		TitleName: "苏克阿赫拉斯",
		TitleCode: "b_soukahras",
	}
}
