package mansura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔曼MirmanBarony struct {
	feud.BaseBarony
}

var BMirman米尔曼 feud.Barony = &米尔曼MirmanBarony{}

func init() {
	f := BMirman米尔曼.(*米尔曼MirmanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirman",
		TitleName: "米尔曼",
		TitleCode: "b_mirman",
	}
}
