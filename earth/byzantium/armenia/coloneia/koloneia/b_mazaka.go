package koloneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马扎卡MazakaBarony struct {
	feud.BaseBarony
}

var BMazaka马扎卡 feud.Barony = &马扎卡MazakaBarony{}

func init() {
	f := BMazaka马扎卡.(*马扎卡MazakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazaka",
		TitleName: "马扎卡",
		TitleCode: "b_mazaka",
	}
}
