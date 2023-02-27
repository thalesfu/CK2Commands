package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马加雷沃MagarevoBarony struct {
	feud.BaseBarony
}

var BMagarevo马加雷沃 feud.Barony = &马加雷沃MagarevoBarony{}

func init() {
    f := BMagarevo马加雷沃.(*马加雷沃MagarevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "magarevo",
		TitleName: "马加雷沃",
		TitleCode: "b_magarevo",
	}
}
