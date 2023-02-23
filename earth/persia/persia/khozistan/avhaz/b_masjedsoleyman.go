package avhaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马斯吉德苏莱曼MasjedsoleymanBarony struct {
	feud.BaseBarony
}

var BMasjedsoleyman马斯吉德苏莱曼 feud.Barony = &马斯吉德苏莱曼MasjedsoleymanBarony{}

func init() {
	f := BMasjedsoleyman马斯吉德苏莱曼.(*马斯吉德苏莱曼MasjedsoleymanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masjedsoleyman",
		TitleName: "马斯吉德苏莱曼",
		TitleCode: "b_masjedsoleyman",
	}
}
