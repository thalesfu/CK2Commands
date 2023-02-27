package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔纳库姆CornacumBarony struct {
	feud.BaseBarony
}

var BCornacum科尔纳库姆 feud.Barony = &科尔纳库姆CornacumBarony{}

func init() {
    f := BCornacum科尔纳库姆.(*科尔纳库姆CornacumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cornacum",
		TitleName: "科尔纳库姆",
		TitleCode: "b_cornacum",
	}
}
