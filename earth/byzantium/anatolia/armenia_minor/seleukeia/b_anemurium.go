package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿内穆里翁AnemuriumBarony struct {
	feud.BaseBarony
}

var BAnemurium阿内穆里翁 feud.Barony = &阿内穆里翁AnemuriumBarony{}

func init() {
    f := BAnemurium阿内穆里翁.(*阿内穆里翁AnemuriumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anemurium",
		TitleName: "阿内穆里翁",
		TitleCode: "b_anemurium",
	}
}
