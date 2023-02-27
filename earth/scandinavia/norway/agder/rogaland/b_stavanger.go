package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯塔万格StavangerBarony struct {
	feud.BaseBarony
}

var BStavanger斯塔万格 feud.Barony = &斯塔万格StavangerBarony{}

func init() {
    f := BStavanger斯塔万格.(*斯塔万格StavangerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stavanger",
		TitleName: "斯塔万格",
		TitleCode: "b_stavanger",
	}
}
