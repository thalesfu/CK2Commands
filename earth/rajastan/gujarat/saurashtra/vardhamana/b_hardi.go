package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曷罗提HardiBarony struct {
	feud.BaseBarony
}

var BHardi曷罗提 feud.Barony = &曷罗提HardiBarony{}

func init() {
	f := BHardi曷罗提.(*曷罗提HardiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hardi",
		TitleName: "曷罗提",
		TitleCode: "b_hardi",
	}
}
