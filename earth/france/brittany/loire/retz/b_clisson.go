package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利松ClissonBarony struct {
	feud.BaseBarony
}

var BClisson克利松 feud.Barony = &克利松ClissonBarony{}

func init() {
    f := BClisson克利松.(*克利松ClissonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clisson",
		TitleName: "克利松",
		TitleCode: "b_clisson",
	}
}
