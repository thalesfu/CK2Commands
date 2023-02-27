package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦泰利VeteliBarony struct {
	feud.BaseBarony
}

var BVeteli韦泰利 feud.Barony = &韦泰利VeteliBarony{}

func init() {
    f := BVeteli韦泰利.(*韦泰利VeteliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veteli",
		TitleName: "韦泰利",
		TitleCode: "b_veteli",
	}
}
