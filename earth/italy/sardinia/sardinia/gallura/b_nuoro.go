package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努奥罗NuoroBarony struct {
	feud.BaseBarony
}

var BNuoro努奥罗 feud.Barony = &努奥罗NuoroBarony{}

func init() {
	f := BNuoro努奥罗.(*努奥罗NuoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nuoro",
		TitleName: "努奥罗",
		TitleCode: "b_nuoro",
	}
}
