package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯于庞厄尔KaupangerBarony struct {
	feud.BaseBarony
}

var BKaupanger凯于庞厄尔 feud.Barony = &凯于庞厄尔KaupangerBarony{}

func init() {
	f := BKaupanger凯于庞厄尔.(*凯于庞厄尔KaupangerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaupanger",
		TitleName: "凯于庞厄尔",
		TitleCode: "b_kaupanger",
	}
}
