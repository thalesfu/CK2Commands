package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅德塔YadetaBarony struct {
	feud.BaseBarony
}

var BYadeta雅德塔 feud.Barony = &雅德塔YadetaBarony{}

func init() {
    f := BYadeta雅德塔.(*雅德塔YadetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yadeta",
		TitleName: "雅德塔",
		TitleCode: "b_yadeta",
	}
}
