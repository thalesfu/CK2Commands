package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达沃蒂亚DawodiyaBarony struct {
	feud.BaseBarony
}

var BDawodiya达沃蒂亚 feud.Barony = &达沃蒂亚DawodiyaBarony{}

func init() {
	f := BDawodiya达沃蒂亚.(*达沃蒂亚DawodiyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dawodiya",
		TitleName: "达沃蒂亚",
		TitleCode: "b_dawodiya",
	}
}
