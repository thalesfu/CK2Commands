package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙特拉永ChatelaillonBarony struct {
	feud.BaseBarony
}

var BChatelaillon沙特拉永 feud.Barony = &沙特拉永ChatelaillonBarony{}

func init() {
    f := BChatelaillon沙特拉永.(*沙特拉永ChatelaillonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatelaillon",
		TitleName: "沙特拉永",
		TitleCode: "b_chatelaillon",
	}
}
