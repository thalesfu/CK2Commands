package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 两姊妹镇DoshermanasBarony struct {
	feud.BaseBarony
}

var BDoshermanas两姊妹镇 feud.Barony = &两姊妹镇DoshermanasBarony{}

func init() {
    f := BDoshermanas两姊妹镇.(*两姊妹镇DoshermanasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doshermanas",
		TitleName: "两姊妹镇",
		TitleCode: "b_doshermanas",
	}
}
