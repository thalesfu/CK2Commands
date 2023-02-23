package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 计剌罗拘吒KerarkotBarony struct {
	feud.BaseBarony
}

var BKerarkot计剌罗拘吒 feud.Barony = &计剌罗拘吒KerarkotBarony{}

func init() {
	f := BKerarkot计剌罗拘吒.(*计剌罗拘吒KerarkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerarkot",
		TitleName: "计剌罗拘吒",
		TitleCode: "b_kerarkot",
	}
}
