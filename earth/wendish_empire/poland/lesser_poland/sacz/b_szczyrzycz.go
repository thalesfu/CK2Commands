package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什哲日茨SzczyrzyczBarony struct {
	feud.BaseBarony
}

var BSzczyrzycz什哲日茨 feud.Barony = &什哲日茨SzczyrzyczBarony{}

func init() {
    f := BSzczyrzycz什哲日茨.(*什哲日茨SzczyrzyczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szczyrzycz",
		TitleName: "什哲日茨",
		TitleCode: "b_szczyrzycz",
	}
}
