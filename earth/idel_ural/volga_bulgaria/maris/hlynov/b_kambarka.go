package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎巴尔卡KambarkaBarony struct {
	feud.BaseBarony
}

var BKambarka坎巴尔卡 feud.Barony = &坎巴尔卡KambarkaBarony{}

func init() {
    f := BKambarka坎巴尔卡.(*坎巴尔卡KambarkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kambarka",
		TitleName: "坎巴尔卡",
		TitleCode: "b_kambarka",
	}
}
