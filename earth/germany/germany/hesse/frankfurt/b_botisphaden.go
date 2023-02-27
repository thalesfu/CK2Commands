package frankfurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博蒂斯法登BotisphadenBarony struct {
	feud.BaseBarony
}

var BBotisphaden博蒂斯法登 feud.Barony = &博蒂斯法登BotisphadenBarony{}

func init() {
    f := BBotisphaden博蒂斯法登.(*博蒂斯法登BotisphadenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "botisphaden",
		TitleName: "博蒂斯法登",
		TitleCode: "b_botisphaden",
	}
}
