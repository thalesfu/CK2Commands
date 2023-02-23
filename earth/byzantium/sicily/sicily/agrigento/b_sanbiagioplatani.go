package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣比亚焦普拉塔尼SanbiagioplataniBarony struct {
	feud.BaseBarony
}

var BSanbiagioplatani圣比亚焦普拉塔尼 feud.Barony = &圣比亚焦普拉塔尼SanbiagioplataniBarony{}

func init() {
	f := BSanbiagioplatani圣比亚焦普拉塔尼.(*圣比亚焦普拉塔尼SanbiagioplataniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanbiagioplatani",
		TitleName: "圣比亚焦普拉塔尼",
		TitleCode: "b_sanbiagioplatani",
	}
}
