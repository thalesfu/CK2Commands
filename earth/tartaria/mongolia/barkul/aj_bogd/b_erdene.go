package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 额尔德尼ErdeneBarony struct {
	feud.BaseBarony
}

var BErdene额尔德尼 feud.Barony = &额尔德尼ErdeneBarony{}

func init() {
    f := BErdene额尔德尼.(*额尔德尼ErdeneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erdene",
		TitleName: "额尔德尼",
		TitleCode: "b_erdene",
	}
}
