package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛萨克SaissacBarony struct {
	feud.BaseBarony
}

var BSaissac赛萨克 feud.Barony = &赛萨克SaissacBarony{}

func init() {
    f := BSaissac赛萨克.(*赛萨克SaissacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saissac",
		TitleName: "赛萨克",
		TitleCode: "b_saissac",
	}
}
