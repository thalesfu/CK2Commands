package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔赫特苏莱曼TakhtesoleymanBarony struct {
	feud.BaseBarony
}

var BTakhtesoleyman塔赫特苏莱曼 feud.Barony = &塔赫特苏莱曼TakhtesoleymanBarony{}

func init() {
	f := BTakhtesoleyman塔赫特苏莱曼.(*塔赫特苏莱曼TakhtesoleymanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takhtesoleyman",
		TitleName: "塔赫特苏莱曼",
		TitleCode: "b_takhtesoleyman",
	}
}
