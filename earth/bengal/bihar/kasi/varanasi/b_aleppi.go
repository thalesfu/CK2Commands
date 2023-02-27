package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚力皮AleppiBarony struct {
	feud.BaseBarony
}

var BAleppi亚力皮 feud.Barony = &亚力皮AleppiBarony{}

func init() {
    f := BAleppi亚力皮.(*亚力皮AleppiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aleppi",
		TitleName: "亚力皮",
		TitleCode: "b_aleppi",
	}
}
