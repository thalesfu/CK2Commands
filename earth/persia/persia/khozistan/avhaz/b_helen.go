package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫伦HelenBarony struct {
	feud.BaseBarony
}

var BHelen赫伦 feud.Barony = &赫伦HelenBarony{}

func init() {
	f := BHelen赫伦.(*赫伦HelenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "helen",
		TitleName: "赫伦",
		TitleCode: "b_helen",
	}
}
