package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特霍姆KastelholmBarony struct {
	feud.BaseBarony
}

var BKastelholm卡斯特霍姆 feud.Barony = &卡斯特霍姆KastelholmBarony{}

func init() {
	f := BKastelholm卡斯特霍姆.(*卡斯特霍姆KastelholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastelholm",
		TitleName: "卡斯特霍姆",
		TitleCode: "b_kastelholm",
	}
}
