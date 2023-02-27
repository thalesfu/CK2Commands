package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希哈克ShihakBarony struct {
	feud.BaseBarony
}

var BShihak希哈克 feud.Barony = &希哈克ShihakBarony{}

func init() {
    f := BShihak希哈克.(*希哈克ShihakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shihak",
		TitleName: "希哈克",
		TitleCode: "b_shihak",
	}
}
