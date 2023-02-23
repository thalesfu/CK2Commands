package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩多耶NdoyeBarony struct {
	feud.BaseBarony
}

var BNdoye恩多耶 feud.Barony = &恩多耶NdoyeBarony{}

func init() {
	f := BNdoye恩多耶.(*恩多耶NdoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ndoye",
		TitleName: "恩多耶",
		TitleCode: "b_ndoye",
	}
}
