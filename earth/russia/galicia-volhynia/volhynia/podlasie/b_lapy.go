package podlasie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦佩LapyBarony struct {
	feud.BaseBarony
}

var BLapy瓦佩 feud.Barony = &瓦佩LapyBarony{}

func init() {
    f := BLapy瓦佩.(*瓦佩LapyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lapy",
		TitleName: "瓦佩",
		TitleCode: "b_lapy",
	}
}
