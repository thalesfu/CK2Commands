package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代莱克DailekhBarony struct {
	feud.BaseBarony
}

var BDailekh代莱克 feud.Barony = &代莱克DailekhBarony{}

func init() {
    f := BDailekh代莱克.(*代莱克DailekhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dailekh",
		TitleName: "代莱克",
		TitleCode: "b_dailekh",
	}
}
