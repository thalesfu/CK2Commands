package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗雷塔VretaBarony struct {
	feud.BaseBarony
}

var BVreta弗雷塔 feud.Barony = &弗雷塔VretaBarony{}

func init() {
	f := BVreta弗雷塔.(*弗雷塔VretaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vreta",
		TitleName: "弗雷塔",
		TitleCode: "b_vreta",
	}
}
