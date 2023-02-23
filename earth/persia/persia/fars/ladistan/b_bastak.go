package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯塔克BastakBarony struct {
	feud.BaseBarony
}

var BBastak巴斯塔克 feud.Barony = &巴斯塔克BastakBarony{}

func init() {
	f := BBastak巴斯塔克.(*巴斯塔克BastakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bastak",
		TitleName: "巴斯塔克",
		TitleCode: "b_bastak",
	}
}
