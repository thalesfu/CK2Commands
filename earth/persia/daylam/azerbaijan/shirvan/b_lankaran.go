package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 连科兰LankaranBarony struct {
	feud.BaseBarony
}

var BLankaran连科兰 feud.Barony = &连科兰LankaranBarony{}

func init() {
	f := BLankaran连科兰.(*连科兰LankaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lankaran",
		TitleName: "连科兰",
		TitleCode: "b_lankaran",
	}
}
