package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊乌扎OlescnicaBarony struct {
	feud.BaseBarony
}

var BOlescnica伊乌扎 feud.Barony = &伊乌扎OlescnicaBarony{}

func init() {
    f := BOlescnica伊乌扎.(*伊乌扎OlescnicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olescnica",
		TitleName: "伊乌扎",
		TitleCode: "b_olescnica",
	}
}
