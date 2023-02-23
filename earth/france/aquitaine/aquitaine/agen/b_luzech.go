package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕泽克LuzechBarony struct {
	feud.BaseBarony
}

var BLuzech吕泽克 feud.Barony = &吕泽克LuzechBarony{}

func init() {
	f := BLuzech吕泽克.(*吕泽克LuzechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luzech",
		TitleName: "吕泽克",
		TitleCode: "b_luzech",
	}
}
