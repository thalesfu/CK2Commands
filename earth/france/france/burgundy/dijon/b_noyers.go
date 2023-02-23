package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努瓦耶NoyersBarony struct {
	feud.BaseBarony
}

var BNoyers努瓦耶 feud.Barony = &努瓦耶NoyersBarony{}

func init() {
	f := BNoyers努瓦耶.(*努瓦耶NoyersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noyers",
		TitleName: "努瓦耶",
		TitleCode: "b_noyers",
	}
}
