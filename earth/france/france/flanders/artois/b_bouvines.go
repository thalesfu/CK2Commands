package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布汶BouvinesBarony struct {
	feud.BaseBarony
}

var BBouvines布汶 feud.Barony = &布汶BouvinesBarony{}

func init() {
    f := BBouvines布汶.(*布汶BouvinesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bouvines",
		TitleName: "布汶",
		TitleCode: "b_bouvines",
	}
}
