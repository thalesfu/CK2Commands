package gwynedd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡那封CaernarfonBarony struct {
	feud.BaseBarony
}

var BCaernarfon卡那封 feud.Barony = &卡那封CaernarfonBarony{}

func init() {
	f := BCaernarfon卡那封.(*卡那封CaernarfonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caernarfon",
		TitleName: "卡那封",
		TitleCode: "b_caernarfon",
	}
}
