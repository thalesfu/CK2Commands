package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯卡HuescaBarony struct {
	feud.BaseBarony
}

var BHuesca韦斯卡 feud.Barony = &韦斯卡HuescaBarony{}

func init() {
    f := BHuesca韦斯卡.(*韦斯卡HuescaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huesca",
		TitleName: "韦斯卡",
		TitleCode: "b_huesca",
	}
}
