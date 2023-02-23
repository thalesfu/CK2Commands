package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普伊PouillyBarony struct {
	feud.BaseBarony
}

var BPouilly普伊 feud.Barony = &普伊PouillyBarony{}

func init() {
	f := BPouilly普伊.(*普伊PouillyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pouilly",
		TitleName: "普伊",
		TitleCode: "b_pouilly",
	}
}
