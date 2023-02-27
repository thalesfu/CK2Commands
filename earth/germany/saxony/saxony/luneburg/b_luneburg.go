package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕讷堡LuneburgBarony struct {
	feud.BaseBarony
}

var BLuneburg吕讷堡 feud.Barony = &吕讷堡LuneburgBarony{}

func init() {
    f := BLuneburg吕讷堡.(*吕讷堡LuneburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luneburg",
		TitleName: "吕讷堡",
		TitleCode: "b_luneburg",
	}
}
