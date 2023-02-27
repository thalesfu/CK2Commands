package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉诺夫VranovBarony struct {
	feud.BaseBarony
}

var BVranov弗拉诺夫 feud.Barony = &弗拉诺夫VranovBarony{}

func init() {
    f := BVranov弗拉诺夫.(*弗拉诺夫VranovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vranov",
		TitleName: "弗拉诺夫",
		TitleCode: "b_vranov",
	}
}
