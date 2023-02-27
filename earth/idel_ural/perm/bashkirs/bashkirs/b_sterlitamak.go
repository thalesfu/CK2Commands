package bashkirs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯捷尔利塔马克SterlitamakBarony struct {
	feud.BaseBarony
}

var BSterlitamak斯捷尔利塔马克 feud.Barony = &斯捷尔利塔马克SterlitamakBarony{}

func init() {
    f := BSterlitamak斯捷尔利塔马克.(*斯捷尔利塔马克SterlitamakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sterlitamak",
		TitleName: "斯捷尔利塔马克",
		TitleCode: "b_sterlitamak",
	}
}
