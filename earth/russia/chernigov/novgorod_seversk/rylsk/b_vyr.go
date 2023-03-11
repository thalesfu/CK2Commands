package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔VyrBarony struct {
	feud.BaseBarony
}

var BVyr维尔 feud.Barony = &维尔VyrBarony{}

func init() {
    f := BVyr维尔.(*维尔VyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyr",
		TitleName: "维尔",
		TitleCode: "b_vyr",
	}
}
