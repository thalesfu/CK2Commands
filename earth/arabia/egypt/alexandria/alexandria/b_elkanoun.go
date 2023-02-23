package alexandria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡努恩ElkanounBarony struct {
	feud.BaseBarony
}

var BElkanoun卡努恩 feud.Barony = &卡努恩ElkanounBarony{}

func init() {
	f := BElkanoun卡努恩.(*卡努恩ElkanounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elkanoun",
		TitleName: "卡努恩",
		TitleCode: "b_elkanoun",
	}
}
