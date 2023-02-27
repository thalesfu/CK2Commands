package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加姆萨尔GarmsarBarony struct {
	feud.BaseBarony
}

var BGarmsar加姆萨尔 feud.Barony = &加姆萨尔GarmsarBarony{}

func init() {
    f := BGarmsar加姆萨尔.(*加姆萨尔GarmsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "garmsar",
		TitleName: "加姆萨尔",
		TitleCode: "b_garmsar",
	}
}
