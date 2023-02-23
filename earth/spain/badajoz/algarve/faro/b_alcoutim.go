package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔科廷AlcoutimBarony struct {
	feud.BaseBarony
}

var BAlcoutim阿尔科廷 feud.Barony = &阿尔科廷AlcoutimBarony{}

func init() {
	f := BAlcoutim阿尔科廷.(*阿尔科廷AlcoutimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcoutim",
		TitleName: "阿尔科廷",
		TitleCode: "b_alcoutim",
	}
}
