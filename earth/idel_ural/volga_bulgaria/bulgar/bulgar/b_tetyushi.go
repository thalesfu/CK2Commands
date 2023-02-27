package bulgar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷秋希TetyushiBarony struct {
	feud.BaseBarony
}

var BTetyushi捷秋希 feud.Barony = &捷秋希TetyushiBarony{}

func init() {
    f := BTetyushi捷秋希.(*捷秋希TetyushiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tetyushi",
		TitleName: "捷秋希",
		TitleCode: "b_tetyushi",
	}
}
