package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼恩NyenBarony struct {
	feud.BaseBarony
}

var BNyen尼恩 feud.Barony = &尼恩NyenBarony{}

func init() {
	f := BNyen尼恩.(*尼恩NyenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyen",
		TitleName: "尼恩",
		TitleCode: "b_nyen",
	}
}
