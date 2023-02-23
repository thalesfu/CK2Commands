package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米林MainlingBarony struct {
	feud.BaseBarony
}

var BMainling米林 feud.Barony = &米林MainlingBarony{}

func init() {
	f := BMainling米林.(*米林MainlingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mainling",
		TitleName: "米林",
		TitleCode: "b_mainling",
	}
}
