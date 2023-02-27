package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大季梅尔卡Velyka_dymerkaBarony struct {
	feud.BaseBarony
}

var BVelyka_dymerka大季梅尔卡 feud.Barony = &大季梅尔卡Velyka_dymerkaBarony{}

func init() {
    f := BVelyka_dymerka大季梅尔卡.(*大季梅尔卡Velyka_dymerkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velyka_dymerka",
		TitleName: "大季梅尔卡",
		TitleCode: "b_velyka_dymerka",
	}
}
