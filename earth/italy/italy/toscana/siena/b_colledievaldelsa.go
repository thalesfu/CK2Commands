package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔萨谷口ColledievaldelsaBarony struct {
	feud.BaseBarony
}

var BColledievaldelsa埃尔萨谷口 feud.Barony = &埃尔萨谷口ColledievaldelsaBarony{}

func init() {
	f := BColledievaldelsa埃尔萨谷口.(*埃尔萨谷口ColledievaldelsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "colledievaldelsa",
		TitleName: "埃尔萨谷口",
		TitleCode: "b_colledievaldelsa",
	}
}
