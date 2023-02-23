package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔斯塔HarstadBarony struct {
	feud.BaseBarony
}

var BHarstad哈尔斯塔 feud.Barony = &哈尔斯塔HarstadBarony{}

func init() {
	f := BHarstad哈尔斯塔.(*哈尔斯塔HarstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harstad",
		TitleName: "哈尔斯塔",
		TitleCode: "b_harstad",
	}
}
