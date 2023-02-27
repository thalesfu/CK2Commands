package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威尼斯VeneziaBarony struct {
	feud.BaseBarony
}

var BVenezia威尼斯 feud.Barony = &威尼斯VeneziaBarony{}

func init() {
    f := BVenezia威尼斯.(*威尼斯VeneziaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "venezia",
		TitleName: "威尼斯",
		TitleCode: "b_venezia",
	}
}
