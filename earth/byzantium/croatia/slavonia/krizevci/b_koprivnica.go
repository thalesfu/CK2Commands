package krizevci

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科普里夫尼察KoprivnicaBarony struct {
	feud.BaseBarony
}

var BKoprivnica科普里夫尼察 feud.Barony = &科普里夫尼察KoprivnicaBarony{}

func init() {
	f := BKoprivnica科普里夫尼察.(*科普里夫尼察KoprivnicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koprivnica",
		TitleName: "科普里夫尼察",
		TitleCode: "b_koprivnica",
	}
}
