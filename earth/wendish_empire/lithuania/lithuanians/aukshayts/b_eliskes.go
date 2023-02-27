package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃利斯克斯EliskesBarony struct {
	feud.BaseBarony
}

var BEliskes埃利斯克斯 feud.Barony = &埃利斯克斯EliskesBarony{}

func init() {
    f := BEliskes埃利斯克斯.(*埃利斯克斯EliskesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eliskes",
		TitleName: "埃利斯克斯",
		TitleCode: "b_eliskes",
	}
}
