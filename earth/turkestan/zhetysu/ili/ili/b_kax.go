package ili

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈什KaxBarony struct {
	feud.BaseBarony
}

var BKax哈什 feud.Barony = &哈什KaxBarony{}

func init() {
	f := BKax哈什.(*哈什KaxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kax",
		TitleName: "哈什",
		TitleCode: "b_kax",
	}
}
