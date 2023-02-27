package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔博MirebeauBarony struct {
	feud.BaseBarony
}

var BMirebeau米尔博 feud.Barony = &米尔博MirebeauBarony{}

func init() {
    f := BMirebeau米尔博.(*米尔博MirebeauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirebeau",
		TitleName: "米尔博",
		TitleCode: "b_mirebeau",
	}
}
