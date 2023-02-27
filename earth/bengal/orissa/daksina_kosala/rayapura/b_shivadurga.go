package rayapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湿婆突伽ShivadurgaBarony struct {
	feud.BaseBarony
}

var BShivadurga湿婆突伽 feud.Barony = &湿婆突伽ShivadurgaBarony{}

func init() {
    f := BShivadurga湿婆突伽.(*湿婆突伽ShivadurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shivadurga",
		TitleName: "湿婆突伽",
		TitleCode: "b_shivadurga",
	}
}
