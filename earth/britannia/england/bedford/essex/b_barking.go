package essex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴金BarkingBarony struct {
	feud.BaseBarony
}

var BBarking巴金 feud.Barony = &巴金BarkingBarony{}

func init() {
    f := BBarking巴金.(*巴金BarkingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barking",
		TitleName: "巴金",
		TitleCode: "b_barking",
	}
}
