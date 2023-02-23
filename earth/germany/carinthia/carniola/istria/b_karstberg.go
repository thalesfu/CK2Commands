package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔斯特堡KarstbergBarony struct {
	feud.BaseBarony
}

var BKarstberg卡尔斯特堡 feud.Barony = &卡尔斯特堡KarstbergBarony{}

func init() {
	f := BKarstberg卡尔斯特堡.(*卡尔斯特堡KarstbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karstberg",
		TitleName: "卡尔斯特堡",
		TitleCode: "b_karstberg",
	}
}
