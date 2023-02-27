package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩尔巴赫MorbaghBarony struct {
	feud.BaseBarony
}

var BMorbagh摩尔巴赫 feud.Barony = &摩尔巴赫MorbaghBarony{}

func init() {
    f := BMorbagh摩尔巴赫.(*摩尔巴赫MorbaghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "morbagh",
		TitleName: "摩尔巴赫",
		TitleCode: "b_morbagh",
	}
}
