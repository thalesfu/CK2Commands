package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代阿斯DeasBarony struct {
	feud.BaseBarony
}

var BDeas代阿斯 feud.Barony = &代阿斯DeasBarony{}

func init() {
    f := BDeas代阿斯.(*代阿斯DeasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deas",
		TitleName: "代阿斯",
		TitleCode: "b_deas",
	}
}
