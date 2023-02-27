package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗塞LandserBarony struct {
	feud.BaseBarony
}

var BLandser朗塞 feud.Barony = &朗塞LandserBarony{}

func init() {
    f := BLandser朗塞.(*朗塞LandserBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "landser",
		TitleName: "朗塞",
		TitleCode: "b_landser",
	}
}
