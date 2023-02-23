package nandana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 候德HundBarony struct {
	feud.BaseBarony
}

var BHund候德 feud.Barony = &候德HundBarony{}

func init() {
	f := BHund候德.(*候德HundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hund",
		TitleName: "候德",
		TitleCode: "b_hund",
	}
}
