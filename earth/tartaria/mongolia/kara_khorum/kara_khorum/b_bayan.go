package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴彦BayanBarony struct {
	feud.BaseBarony
}

var BBayan巴彦 feud.Barony = &巴彦BayanBarony{}

func init() {
    f := BBayan巴彦.(*巴彦BayanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayan",
		TitleName: "巴彦",
		TitleCode: "b_bayan",
	}
}
