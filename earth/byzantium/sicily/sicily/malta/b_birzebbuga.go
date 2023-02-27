package malta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇塔代拉BirzebbugaBarony struct {
	feud.BaseBarony
}

var BBirzebbuga奇塔代拉 feud.Barony = &奇塔代拉BirzebbugaBarony{}

func init() {
    f := BBirzebbuga奇塔代拉.(*奇塔代拉BirzebbugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birzebbuga",
		TitleName: "奇塔代拉",
		TitleCode: "b_birzebbuga",
	}
}
