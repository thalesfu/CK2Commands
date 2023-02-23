package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟尔斯坦ThirlestaneBarony struct {
	feud.BaseBarony
}

var BThirlestane瑟尔斯坦 feud.Barony = &瑟尔斯坦ThirlestaneBarony{}

func init() {
	f := BThirlestane瑟尔斯坦.(*瑟尔斯坦ThirlestaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thirlestane",
		TitleName: "瑟尔斯坦",
		TitleCode: "b_thirlestane",
	}
}
