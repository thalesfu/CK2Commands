package azerbaijan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔赫特苏莱曼TakhtesuleimanBarony struct {
	feud.BaseBarony
}

var BTakhtesuleiman塔赫特苏莱曼 feud.Barony = &塔赫特苏莱曼TakhtesuleimanBarony{}

func init() {
    f := BTakhtesuleiman塔赫特苏莱曼.(*塔赫特苏莱曼TakhtesuleimanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takhtesuleiman",
		TitleName: "塔赫特苏莱曼",
		TitleCode: "b_takhtesuleiman",
	}
}
