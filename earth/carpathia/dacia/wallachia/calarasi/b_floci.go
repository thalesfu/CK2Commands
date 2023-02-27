package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗洛奇FlociBarony struct {
	feud.BaseBarony
}

var BFloci弗洛奇 feud.Barony = &弗洛奇FlociBarony{}

func init() {
    f := BFloci弗洛奇.(*弗洛奇FlociBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "floci",
		TitleName: "弗洛奇",
		TitleCode: "b_floci",
	}
}
