package sarpa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克切涅雷KetcheneryBarony struct {
	feud.BaseBarony
}

var BKetchenery克切涅雷 feud.Barony = &克切涅雷KetcheneryBarony{}

func init() {
    f := BKetchenery克切涅雷.(*克切涅雷KetcheneryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ketchenery",
		TitleName: "克切涅雷",
		TitleCode: "b_ketchenery",
	}
}
