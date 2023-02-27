package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗斯诺KrosnoBarony struct {
	feud.BaseBarony
}

var BKrosno克罗斯诺 feud.Barony = &克罗斯诺KrosnoBarony{}

func init() {
    f := BKrosno克罗斯诺.(*克罗斯诺KrosnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krosno",
		TitleName: "克罗斯诺",
		TitleCode: "b_krosno",
	}
}
