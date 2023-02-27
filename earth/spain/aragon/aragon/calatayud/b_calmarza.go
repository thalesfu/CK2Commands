package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔马尔萨CalmarzaBarony struct {
	feud.BaseBarony
}

var BCalmarza卡尔马尔萨 feud.Barony = &卡尔马尔萨CalmarzaBarony{}

func init() {
    f := BCalmarza卡尔马尔萨.(*卡尔马尔萨CalmarzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calmarza",
		TitleName: "卡尔马尔萨",
		TitleCode: "b_calmarza",
	}
}
