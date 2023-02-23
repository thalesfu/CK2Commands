package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索波特SopotBarony struct {
	feud.BaseBarony
}

var BSopot索波特 feud.Barony = &索波特SopotBarony{}

func init() {
	f := BSopot索波特.(*索波特SopotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sopot",
		TitleName: "索波特",
		TitleCode: "b_sopot",
	}
}
