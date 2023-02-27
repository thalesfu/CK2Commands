package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳迪基亚Laodicea_sinopeBarony struct {
	feud.BaseBarony
}

var BLaodicea_sinope劳迪基亚 feud.Barony = &劳迪基亚Laodicea_sinopeBarony{}

func init() {
    f := BLaodicea_sinope劳迪基亚.(*劳迪基亚Laodicea_sinopeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laodicea_sinope",
		TitleName: "劳迪基亚",
		TitleCode: "b_laodicea_sinope",
	}
}
