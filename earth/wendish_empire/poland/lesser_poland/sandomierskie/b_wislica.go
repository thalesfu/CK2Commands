package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维希利察WislicaBarony struct {
	feud.BaseBarony
}

var BWislica维希利察 feud.Barony = &维希利察WislicaBarony{}

func init() {
    f := BWislica维希利察.(*维希利察WislicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wislica",
		TitleName: "维希利察",
		TitleCode: "b_wislica",
	}
}
