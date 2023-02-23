package aulon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯皮纳里扎SpinarizzaBarony struct {
	feud.BaseBarony
}

var BSpinarizza斯皮纳里扎 feud.Barony = &斯皮纳里扎SpinarizzaBarony{}

func init() {
	f := BSpinarizza斯皮纳里扎.(*斯皮纳里扎SpinarizzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spinarizza",
		TitleName: "斯皮纳里扎",
		TitleCode: "b_spinarizza",
	}
}
