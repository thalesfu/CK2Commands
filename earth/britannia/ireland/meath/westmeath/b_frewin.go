package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗雷温FrewinBarony struct {
	feud.BaseBarony
}

var BFrewin弗雷温 feud.Barony = &弗雷温FrewinBarony{}

func init() {
    f := BFrewin弗雷温.(*弗雷温FrewinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frewin",
		TitleName: "弗雷温",
		TitleCode: "b_frewin",
	}
}
