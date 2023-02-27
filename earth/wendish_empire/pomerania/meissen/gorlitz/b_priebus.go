package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里布斯PriebusBarony struct {
	feud.BaseBarony
}

var BPriebus普里布斯 feud.Barony = &普里布斯PriebusBarony{}

func init() {
    f := BPriebus普里布斯.(*普里布斯PriebusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "priebus",
		TitleName: "普里布斯",
		TitleCode: "b_priebus",
	}
}
