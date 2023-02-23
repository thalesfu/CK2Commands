package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波热加PozegaBarony struct {
	feud.BaseBarony
}

var BPozega波热加 feud.Barony = &波热加PozegaBarony{}

func init() {
	f := BPozega波热加.(*波热加PozegaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pozega",
		TitleName: "波热加",
		TitleCode: "b_pozega",
	}
}
