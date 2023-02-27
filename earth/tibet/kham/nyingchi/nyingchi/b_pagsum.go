package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 八松PagsumBarony struct {
	feud.BaseBarony
}

var BPagsum八松 feud.Barony = &八松PagsumBarony{}

func init() {
    f := BPagsum八松.(*八松PagsumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pagsum",
		TitleName: "八松",
		TitleCode: "b_pagsum",
	}
}
