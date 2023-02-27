package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔法利亚TafallaBarony struct {
	feud.BaseBarony
}

var BTafalla塔法利亚 feud.Barony = &塔法利亚TafallaBarony{}

func init() {
    f := BTafalla塔法利亚.(*塔法利亚TafallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tafalla",
		TitleName: "塔法利亚",
		TitleCode: "b_tafalla",
	}
}
