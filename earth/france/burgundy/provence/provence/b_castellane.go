package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰朗CastellaneBarony struct {
	feud.BaseBarony
}

var BCastellane卡斯泰朗 feud.Barony = &卡斯泰朗CastellaneBarony{}

func init() {
    f := BCastellane卡斯泰朗.(*卡斯泰朗CastellaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castellane",
		TitleName: "卡斯泰朗",
		TitleCode: "b_castellane",
	}
}
