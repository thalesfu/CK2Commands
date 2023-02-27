package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 定结DinggyeBarony struct {
	feud.BaseBarony
}

var BDinggye定结 feud.Barony = &定结DinggyeBarony{}

func init() {
    f := BDinggye定结.(*定结DinggyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dinggye",
		TitleName: "定结",
		TitleCode: "b_dinggye",
	}
}
