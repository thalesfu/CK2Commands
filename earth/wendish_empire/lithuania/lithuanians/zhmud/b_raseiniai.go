package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉塞尼艾RaseiniaiBarony struct {
	feud.BaseBarony
}

var BRaseiniai拉塞尼艾 feud.Barony = &拉塞尼艾RaseiniaiBarony{}

func init() {
    f := BRaseiniai拉塞尼艾.(*拉塞尼艾RaseiniaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raseiniai",
		TitleName: "拉塞尼艾",
		TitleCode: "b_raseiniai",
	}
}
