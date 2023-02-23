package tangiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆莱布瑟罕MulaybuselhamBarony struct {
	feud.BaseBarony
}

var BMulaybuselham穆莱布瑟罕 feud.Barony = &穆莱布瑟罕MulaybuselhamBarony{}

func init() {
	f := BMulaybuselham穆莱布瑟罕.(*穆莱布瑟罕MulaybuselhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mulaybuselham",
		TitleName: "穆莱布瑟罕",
		TitleCode: "b_mulaybuselham",
	}
}
