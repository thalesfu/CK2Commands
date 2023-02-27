package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈罗杰茨GorodezBarony struct {
	feud.BaseBarony
}

var BGorodez戈罗杰茨 feud.Barony = &戈罗杰茨GorodezBarony{}

func init() {
    f := BGorodez戈罗杰茨.(*戈罗杰茨GorodezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gorodez",
		TitleName: "戈罗杰茨",
		TitleCode: "b_gorodez",
	}
}
