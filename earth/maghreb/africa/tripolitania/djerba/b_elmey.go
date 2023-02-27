package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马伊ElmeyBarony struct {
	feud.BaseBarony
}

var BElmey马伊 feud.Barony = &马伊ElmeyBarony{}

func init() {
    f := BElmey马伊.(*马伊ElmeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elmey",
		TitleName: "马伊",
		TitleCode: "b_elmey",
	}
}
