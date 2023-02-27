package khuiten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇赫尔特ChikherteiBarony struct {
	feud.BaseBarony
}

var BChikhertei奇赫尔特 feud.Barony = &奇赫尔特ChikherteiBarony{}

func init() {
    f := BChikhertei奇赫尔特.(*奇赫尔特ChikherteiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chikhertei",
		TitleName: "奇赫尔特",
		TitleCode: "b_chikhertei",
	}
}
