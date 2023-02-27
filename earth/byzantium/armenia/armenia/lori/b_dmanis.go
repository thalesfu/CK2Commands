package lori

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德曼尼斯DmanisBarony struct {
	feud.BaseBarony
}

var BDmanis德曼尼斯 feud.Barony = &德曼尼斯DmanisBarony{}

func init() {
    f := BDmanis德曼尼斯.(*德曼尼斯DmanisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dmanis",
		TitleName: "德曼尼斯",
		TitleCode: "b_dmanis",
	}
}
