package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赞布鲁夫ZambrowBarony struct {
	feud.BaseBarony
}

var BZambrow赞布鲁夫 feud.Barony = &赞布鲁夫ZambrowBarony{}

func init() {
    f := BZambrow赞布鲁夫.(*赞布鲁夫ZambrowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zambrow",
		TitleName: "赞布鲁夫",
		TitleCode: "b_zambrow",
	}
}
