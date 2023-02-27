package upper_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗罗茨瓦夫WroclawBarony struct {
	feud.BaseBarony
}

var BWroclaw弗罗茨瓦夫 feud.Barony = &弗罗茨瓦夫WroclawBarony{}

func init() {
    f := BWroclaw弗罗茨瓦夫.(*弗罗茨瓦夫WroclawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wroclaw",
		TitleName: "弗罗茨瓦夫",
		TitleCode: "b_wroclaw",
	}
}
