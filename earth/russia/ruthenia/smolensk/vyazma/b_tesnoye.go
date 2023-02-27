package vyazma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷斯诺耶TesnoyeBarony struct {
	feud.BaseBarony
}

var BTesnoye捷斯诺耶 feud.Barony = &捷斯诺耶TesnoyeBarony{}

func init() {
    f := BTesnoye捷斯诺耶.(*捷斯诺耶TesnoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tesnoye",
		TitleName: "捷斯诺耶",
		TitleCode: "b_tesnoye",
	}
}
