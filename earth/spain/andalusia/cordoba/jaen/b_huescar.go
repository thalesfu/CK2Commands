package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯卡尔HuescarBarony struct {
	feud.BaseBarony
}

var BHuescar韦斯卡尔 feud.Barony = &韦斯卡尔HuescarBarony{}

func init() {
	f := BHuescar韦斯卡尔.(*韦斯卡尔HuescarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huescar",
		TitleName: "韦斯卡尔",
		TitleCode: "b_huescar",
	}
}
