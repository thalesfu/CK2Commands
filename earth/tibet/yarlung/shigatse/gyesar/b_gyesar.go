package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰萨GyesarBarony struct {
	feud.BaseBarony
}

var BGyesar杰萨 feud.Barony = &杰萨GyesarBarony{}

func init() {
	f := BGyesar杰萨.(*杰萨GyesarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyesar",
		TitleName: "杰萨",
		TitleCode: "b_gyesar",
	}
}
