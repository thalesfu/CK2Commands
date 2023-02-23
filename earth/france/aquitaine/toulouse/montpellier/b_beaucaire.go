package montpellier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博凯尔BeaucaireBarony struct {
	feud.BaseBarony
}

var BBeaucaire博凯尔 feud.Barony = &博凯尔BeaucaireBarony{}

func init() {
	f := BBeaucaire博凯尔.(*博凯尔BeaucaireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaucaire",
		TitleName: "博凯尔",
		TitleCode: "b_beaucaire",
	}
}
