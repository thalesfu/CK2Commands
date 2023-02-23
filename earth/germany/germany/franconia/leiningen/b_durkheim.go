package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪尔凯姆DurkheimBarony struct {
	feud.BaseBarony
}

var BDurkheim迪尔凯姆 feud.Barony = &迪尔凯姆DurkheimBarony{}

func init() {
	f := BDurkheim迪尔凯姆.(*迪尔凯姆DurkheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durkheim",
		TitleName: "迪尔凯姆",
		TitleCode: "b_durkheim",
	}
}
