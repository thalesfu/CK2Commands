package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂齐乌祖TiziouzouBarony struct {
	feud.BaseBarony
}

var BTiziouzou蒂齐乌祖 feud.Barony = &蒂齐乌祖TiziouzouBarony{}

func init() {
    f := BTiziouzou蒂齐乌祖.(*蒂齐乌祖TiziouzouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiziouzou",
		TitleName: "蒂齐乌祖",
		TitleCode: "b_tiziouzou",
	}
}
