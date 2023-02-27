package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 金塔纳尔德尔雷伊QuintanardelreyBarony struct {
	feud.BaseBarony
}

var BQuintanardelrey金塔纳尔德尔雷伊 feud.Barony = &金塔纳尔德尔雷伊QuintanardelreyBarony{}

func init() {
    f := BQuintanardelrey金塔纳尔德尔雷伊.(*金塔纳尔德尔雷伊QuintanardelreyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quintanardelrey",
		TitleName: "金塔纳尔德尔雷伊",
		TitleCode: "b_quintanardelrey",
	}
}
