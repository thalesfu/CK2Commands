package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塞格KasegBarony struct {
	feud.BaseBarony
}

var BKaseg卡塞格 feud.Barony = &卡塞格KasegBarony{}

func init() {
    f := BKaseg卡塞格.(*卡塞格KasegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaseg",
		TitleName: "卡塞格",
		TitleCode: "b_kaseg",
	}
}
