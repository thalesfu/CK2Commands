package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯纳瓦斯德尔马德罗尼奥LasnavasdelmadronoBarony struct {
	feud.BaseBarony
}

var BLasnavasdelmadrono拉斯纳瓦斯德尔马德罗尼奥 feud.Barony = &拉斯纳瓦斯德尔马德罗尼奥LasnavasdelmadronoBarony{}

func init() {
	f := BLasnavasdelmadrono拉斯纳瓦斯德尔马德罗尼奥.(*拉斯纳瓦斯德尔马德罗尼奥LasnavasdelmadronoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lasnavasdelmadrono",
		TitleName: "拉斯纳瓦斯德尔马德罗尼奥",
		TitleCode: "b_lasnavasdelmadrono",
	}
}
