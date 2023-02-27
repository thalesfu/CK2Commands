package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆巴拉斯AmbarasBarony struct {
	feud.BaseBarony
}

var BAmbaras阿姆巴拉斯 feud.Barony = &阿姆巴拉斯AmbarasBarony{}

func init() {
    f := BAmbaras阿姆巴拉斯.(*阿姆巴拉斯AmbarasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ambaras",
		TitleName: "阿姆巴拉斯",
		TitleCode: "b_ambaras",
	}
}
