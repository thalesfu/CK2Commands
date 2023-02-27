package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湿婆城ShivapuraBarony struct {
	feud.BaseBarony
}

var BShivapura湿婆城 feud.Barony = &湿婆城ShivapuraBarony{}

func init() {
    f := BShivapura湿婆城.(*湿婆城ShivapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shivapura",
		TitleName: "湿婆城",
		TitleCode: "b_shivapura",
	}
}
