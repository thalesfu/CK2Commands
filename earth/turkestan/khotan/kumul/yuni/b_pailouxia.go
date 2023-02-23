package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牌楼下PailouxiaBarony struct {
	feud.BaseBarony
}

var BPailouxia牌楼下 feud.Barony = &牌楼下PailouxiaBarony{}

func init() {
	f := BPailouxia牌楼下.(*牌楼下PailouxiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pailouxia",
		TitleName: "牌楼下",
		TitleCode: "b_pailouxia",
	}
}
