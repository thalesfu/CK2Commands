package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法锡斯PhasisBarony struct {
	feud.BaseBarony
}

var BPhasis法锡斯 feud.Barony = &法锡斯PhasisBarony{}

func init() {
    f := BPhasis法锡斯.(*法锡斯PhasisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phasis",
		TitleName: "法锡斯",
		TitleCode: "b_phasis",
	}
}
