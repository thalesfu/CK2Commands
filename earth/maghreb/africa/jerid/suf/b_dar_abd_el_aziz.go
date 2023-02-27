package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔阿卜杜勒阿齐兹Dar_abd_el_azizBarony struct {
	feud.BaseBarony
}

var BDar_abd_el_aziz达尔阿卜杜勒阿齐兹 feud.Barony = &达尔阿卜杜勒阿齐兹Dar_abd_el_azizBarony{}

func init() {
    f := BDar_abd_el_aziz达尔阿卜杜勒阿齐兹.(*达尔阿卜杜勒阿齐兹Dar_abd_el_azizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dar_abd_el_aziz",
		TitleName: "达尔阿卜杜勒阿齐兹",
		TitleCode: "b_dar_abd_el_aziz",
	}
}
