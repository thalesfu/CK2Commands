package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安斯巴赫AnsbachBarony struct {
	feud.BaseBarony
}

var BAnsbach安斯巴赫 feud.Barony = &安斯巴赫AnsbachBarony{}

func init() {
	f := BAnsbach安斯巴赫.(*安斯巴赫AnsbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ansbach",
		TitleName: "安斯巴赫",
		TitleCode: "b_ansbach",
	}
}
