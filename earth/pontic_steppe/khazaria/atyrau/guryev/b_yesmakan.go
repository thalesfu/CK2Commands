package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶斯马坎YesmakanBarony struct {
	feud.BaseBarony
}

var BYesmakan叶斯马坎 feud.Barony = &叶斯马坎YesmakanBarony{}

func init() {
    f := BYesmakan叶斯马坎.(*叶斯马坎YesmakanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yesmakan",
		TitleName: "叶斯马坎",
		TitleCode: "b_yesmakan",
	}
}
