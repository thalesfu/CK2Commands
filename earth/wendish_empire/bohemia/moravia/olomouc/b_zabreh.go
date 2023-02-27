package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎布热赫ZabrehBarony struct {
	feud.BaseBarony
}

var BZabreh扎布热赫 feud.Barony = &扎布热赫ZabrehBarony{}

func init() {
    f := BZabreh扎布热赫.(*扎布热赫ZabrehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zabreh",
		TitleName: "扎布热赫",
		TitleCode: "b_zabreh",
	}
}
