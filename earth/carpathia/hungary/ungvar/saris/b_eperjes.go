package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾派尔耶什EperjesBarony struct {
	feud.BaseBarony
}

var BEperjes艾派尔耶什 feud.Barony = &艾派尔耶什EperjesBarony{}

func init() {
	f := BEperjes艾派尔耶什.(*艾派尔耶什EperjesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eperjes",
		TitleName: "艾派尔耶什",
		TitleCode: "b_eperjes",
	}
}
