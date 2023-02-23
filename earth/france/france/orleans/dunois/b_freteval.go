package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗雷特瓦勒FretevalBarony struct {
	feud.BaseBarony
}

var BFreteval弗雷特瓦勒 feud.Barony = &弗雷特瓦勒FretevalBarony{}

func init() {
	f := BFreteval弗雷特瓦勒.(*弗雷特瓦勒FretevalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "freteval",
		TitleName: "弗雷特瓦勒",
		TitleCode: "b_freteval",
	}
}
