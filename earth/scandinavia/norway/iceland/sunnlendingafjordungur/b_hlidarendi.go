package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫利达伦迪HlidarendiBarony struct {
	feud.BaseBarony
}

var BHlidarendi赫利达伦迪 feud.Barony = &赫利达伦迪HlidarendiBarony{}

func init() {
    f := BHlidarendi赫利达伦迪.(*赫利达伦迪HlidarendiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hlidarendi",
		TitleName: "赫利达伦迪",
		TitleCode: "b_hlidarendi",
	}
}
