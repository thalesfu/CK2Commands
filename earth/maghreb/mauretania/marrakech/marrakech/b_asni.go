package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾斯尼AsniBarony struct {
	feud.BaseBarony
}

var BAsni艾斯尼 feud.Barony = &艾斯尼AsniBarony{}

func init() {
	f := BAsni艾斯尼.(*艾斯尼AsniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asni",
		TitleName: "艾斯尼",
		TitleCode: "b_asni",
	}
}
