package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普热罗夫PrerovBarony struct {
	feud.BaseBarony
}

var BPrerov普热罗夫 feud.Barony = &普热罗夫PrerovBarony{}

func init() {
    f := BPrerov普热罗夫.(*普热罗夫PrerovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prerov",
		TitleName: "普热罗夫",
		TitleCode: "b_prerov",
	}
}
