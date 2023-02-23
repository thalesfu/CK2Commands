package dipalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇查瓦尼ChichawatniBarony struct {
	feud.BaseBarony
}

var BChichawatni奇查瓦尼 feud.Barony = &奇查瓦尼ChichawatniBarony{}

func init() {
	f := BChichawatni奇查瓦尼.(*奇查瓦尼ChichawatniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chichawatni",
		TitleName: "奇查瓦尼",
		TitleCode: "b_chichawatni",
	}
}
