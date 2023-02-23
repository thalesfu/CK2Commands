package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布斯卡BuscaBarony struct {
	feud.BaseBarony
}

var BBusca布斯卡 feud.Barony = &布斯卡BuscaBarony{}

func init() {
	f := BBusca布斯卡.(*布斯卡BuscaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "busca",
		TitleName: "布斯卡",
		TitleCode: "b_busca",
	}
}
