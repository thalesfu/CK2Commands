package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼尸HurshiBarony struct {
	feud.BaseBarony
}

var BHurshi呼尸 feud.Barony = &呼尸HurshiBarony{}

func init() {
	f := BHurshi呼尸.(*呼尸HurshiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hurshi",
		TitleName: "呼尸",
		TitleCode: "b_hurshi",
	}
}
