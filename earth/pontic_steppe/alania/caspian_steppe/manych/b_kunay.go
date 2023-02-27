package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库奈KunayBarony struct {
	feud.BaseBarony
}

var BKunay库奈 feud.Barony = &库奈KunayBarony{}

func init() {
    f := BKunay库奈.(*库奈KunayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunay",
		TitleName: "库奈",
		TitleCode: "b_kunay",
	}
}
