package dege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 八邦PalpungBarony struct {
	feud.BaseBarony
}

var BPalpung八邦 feud.Barony = &八邦PalpungBarony{}

func init() {
	f := BPalpung八邦.(*八邦PalpungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palpung",
		TitleName: "八邦",
		TitleCode: "b_palpung",
	}
}
