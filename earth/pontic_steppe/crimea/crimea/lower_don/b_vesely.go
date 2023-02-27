package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦谢雷VeselyBarony struct {
	feud.BaseBarony
}

var BVesely韦谢雷 feud.Barony = &韦谢雷VeselyBarony{}

func init() {
    f := BVesely韦谢雷.(*韦谢雷VeselyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vesely",
		TitleName: "韦谢雷",
		TitleCode: "b_vesely",
	}
}
