package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克塞格KoszegBarony struct {
	feud.BaseBarony
}

var BKoszeg克塞格 feud.Barony = &克塞格KoszegBarony{}

func init() {
    f := BKoszeg克塞格.(*克塞格KoszegBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koszeg",
		TitleName: "克塞格",
		TitleCode: "b_koszeg",
	}
}
