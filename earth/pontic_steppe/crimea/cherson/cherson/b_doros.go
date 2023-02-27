package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗斯DorosBarony struct {
	feud.BaseBarony
}

var BDoros多罗斯 feud.Barony = &多罗斯DorosBarony{}

func init() {
    f := BDoros多罗斯.(*多罗斯DorosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doros",
		TitleName: "多罗斯",
		TitleCode: "b_doros",
	}
}
