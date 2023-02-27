package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍多宁HodoninBarony struct {
	feud.BaseBarony
}

var BHodonin霍多宁 feud.Barony = &霍多宁HodoninBarony{}

func init() {
    f := BHodonin霍多宁.(*霍多宁HodoninBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hodonin",
		TitleName: "霍多宁",
		TitleCode: "b_hodonin",
	}
}
