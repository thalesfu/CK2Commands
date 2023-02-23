package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布托夫卡ButovkaBarony struct {
	feud.BaseBarony
}

var BButovka布托夫卡 feud.Barony = &布托夫卡ButovkaBarony{}

func init() {
	f := BButovka布托夫卡.(*布托夫卡ButovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "butovka",
		TitleName: "布托夫卡",
		TitleCode: "b_butovka",
	}
}
