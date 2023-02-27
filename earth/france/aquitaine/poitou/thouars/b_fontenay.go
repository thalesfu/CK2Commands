package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丰特奈FontenayBarony struct {
	feud.BaseBarony
}

var BFontenay丰特奈 feud.Barony = &丰特奈FontenayBarony{}

func init() {
    f := BFontenay丰特奈.(*丰特奈FontenayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fontenay",
		TitleName: "丰特奈",
		TitleCode: "b_fontenay",
	}
}
