package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济咙GyirongBarony struct {
	feud.BaseBarony
}

var BGyirong济咙 feud.Barony = &济咙GyirongBarony{}

func init() {
    f := BGyirong济咙.(*济咙GyirongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyirong",
		TitleName: "济咙",
		TitleCode: "b_gyirong",
	}
}
