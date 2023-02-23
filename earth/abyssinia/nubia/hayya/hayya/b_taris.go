package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔里斯TarisBarony struct {
	feud.BaseBarony
}

var BTaris塔里斯 feud.Barony = &塔里斯TarisBarony{}

func init() {
	f := BTaris塔里斯.(*塔里斯TarisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taris",
		TitleName: "塔里斯",
		TitleCode: "b_taris",
	}
}
