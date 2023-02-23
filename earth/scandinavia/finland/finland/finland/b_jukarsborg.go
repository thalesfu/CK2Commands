package finland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于卡斯伯格JukarsborgBarony struct {
	feud.BaseBarony
}

var BJukarsborg于卡斯伯格 feud.Barony = &于卡斯伯格JukarsborgBarony{}

func init() {
	f := BJukarsborg于卡斯伯格.(*于卡斯伯格JukarsborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jukarsborg",
		TitleName: "于卡斯伯格",
		TitleCode: "b_jukarsborg",
	}
}
