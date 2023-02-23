package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗雷瑞斯FrejusBarony struct {
	feud.BaseBarony
}

var BFrejus弗雷瑞斯 feud.Barony = &弗雷瑞斯FrejusBarony{}

func init() {
	f := BFrejus弗雷瑞斯.(*弗雷瑞斯FrejusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frejus",
		TitleName: "弗雷瑞斯",
		TitleCode: "b_frejus",
	}
}
