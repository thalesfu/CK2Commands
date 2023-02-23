package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴哈雷AlbaharahBarony struct {
	feud.BaseBarony
}

var BAlbaharah巴哈雷 feud.Barony = &巴哈雷AlbaharahBarony{}

func init() {
	f := BAlbaharah巴哈雷.(*巴哈雷AlbaharahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albaharah",
		TitleName: "巴哈雷",
		TitleCode: "b_albaharah",
	}
}
