package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙力NauliBarony struct {
	feud.BaseBarony
}

var BNauli瑙力 feud.Barony = &瑙力NauliBarony{}

func init() {
    f := BNauli瑙力.(*瑙力NauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nauli",
		TitleName: "瑙力",
		TitleCode: "b_nauli",
	}
}
