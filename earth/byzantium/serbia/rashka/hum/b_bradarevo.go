package hum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布罗达雷沃BradarevoBarony struct {
	feud.BaseBarony
}

var BBradarevo布罗达雷沃 feud.Barony = &布罗达雷沃BradarevoBarony{}

func init() {
    f := BBradarevo布罗达雷沃.(*布罗达雷沃BradarevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bradarevo",
		TitleName: "布罗达雷沃",
		TitleCode: "b_bradarevo",
	}
}
