package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维耶斯泰ViesteBarony struct {
	feud.BaseBarony
}

var BVieste维耶斯泰 feud.Barony = &维耶斯泰ViesteBarony{}

func init() {
    f := BVieste维耶斯泰.(*维耶斯泰ViesteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vieste",
		TitleName: "维耶斯泰",
		TitleCode: "b_vieste",
	}
}
