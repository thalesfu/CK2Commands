package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡盖特堡SzigetvarBarony struct {
	feud.BaseBarony
}

var BSzigetvar锡盖特堡 feud.Barony = &锡盖特堡SzigetvarBarony{}

func init() {
	f := BSzigetvar锡盖特堡.(*锡盖特堡SzigetvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szigetvar",
		TitleName: "锡盖特堡",
		TitleCode: "b_szigetvar",
	}
}
