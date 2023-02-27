package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊诺弗罗茨瓦夫InowroclawBarony struct {
	feud.BaseBarony
}

var BInowroclaw伊诺弗罗茨瓦夫 feud.Barony = &伊诺弗罗茨瓦夫InowroclawBarony{}

func init() {
    f := BInowroclaw伊诺弗罗茨瓦夫.(*伊诺弗罗茨瓦夫InowroclawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inowroclaw",
		TitleName: "伊诺弗罗茨瓦夫",
		TitleCode: "b_inowroclaw",
	}
}
