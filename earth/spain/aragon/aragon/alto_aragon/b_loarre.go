package alto_aragon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛阿雷LoarreBarony struct {
	feud.BaseBarony
}

var BLoarre洛阿雷 feud.Barony = &洛阿雷LoarreBarony{}

func init() {
    f := BLoarre洛阿雷.(*洛阿雷LoarreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loarre",
		TitleName: "洛阿雷",
		TitleCode: "b_loarre",
	}
}
