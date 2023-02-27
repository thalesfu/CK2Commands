package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁瓦新城VilleneuveleroiBarony struct {
	feud.BaseBarony
}

var BVilleneuveleroi鲁瓦新城 feud.Barony = &鲁瓦新城VilleneuveleroiBarony{}

func init() {
    f := BVilleneuveleroi鲁瓦新城.(*鲁瓦新城VilleneuveleroiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "villeneuveleroi",
		TitleName: "鲁瓦新城",
		TitleCode: "b_villeneuveleroi",
	}
}
