package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃利特莱ErythraiBarony struct {
	feud.BaseBarony
}

var BErythrai埃利特莱 feud.Barony = &埃利特莱ErythraiBarony{}

func init() {
    f := BErythrai埃利特莱.(*埃利特莱ErythraiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erythrai",
		TitleName: "埃利特莱",
		TitleCode: "b_erythrai",
	}
}
