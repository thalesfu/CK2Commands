package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯塔巴拉CastabalaBarony struct {
	feud.BaseBarony
}

var BCastabala卡斯塔巴拉 feud.Barony = &卡斯塔巴拉CastabalaBarony{}

func init() {
    f := BCastabala卡斯塔巴拉.(*卡斯塔巴拉CastabalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castabala",
		TitleName: "卡斯塔巴拉",
		TitleCode: "b_castabala",
	}
}
