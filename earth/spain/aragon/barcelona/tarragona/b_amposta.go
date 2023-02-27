package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安波斯塔AmpostaBarony struct {
	feud.BaseBarony
}

var BAmposta安波斯塔 feud.Barony = &安波斯塔AmpostaBarony{}

func init() {
    f := BAmposta安波斯塔.(*安波斯塔AmpostaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amposta",
		TitleName: "安波斯塔",
		TitleCode: "b_amposta",
	}
}
