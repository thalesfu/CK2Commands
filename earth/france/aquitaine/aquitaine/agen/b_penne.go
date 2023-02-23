package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩纳PenneBarony struct {
	feud.BaseBarony
}

var BPenne佩纳 feud.Barony = &佩纳PenneBarony{}

func init() {
	f := BPenne佩纳.(*佩纳PenneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penne",
		TitleName: "佩纳",
		TitleCode: "b_penne",
	}
}
