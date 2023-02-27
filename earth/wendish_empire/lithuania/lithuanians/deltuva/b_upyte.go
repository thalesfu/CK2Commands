package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌皮特UpyteBarony struct {
	feud.BaseBarony
}

var BUpyte乌皮特 feud.Barony = &乌皮特UpyteBarony{}

func init() {
    f := BUpyte乌皮特.(*乌皮特UpyteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "upyte",
		TitleName: "乌皮特",
		TitleCode: "b_upyte",
	}
}
