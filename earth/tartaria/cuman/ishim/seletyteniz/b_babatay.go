package seletyteniz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴巴泰BabatayBarony struct {
	feud.BaseBarony
}

var BBabatay巴巴泰 feud.Barony = &巴巴泰BabatayBarony{}

func init() {
    f := BBabatay巴巴泰.(*巴巴泰BabatayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babatay",
		TitleName: "巴巴泰",
		TitleCode: "b_babatay",
	}
}
