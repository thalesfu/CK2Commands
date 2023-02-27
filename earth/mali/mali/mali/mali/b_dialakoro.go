package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪亚拉科罗DialakoroBarony struct {
	feud.BaseBarony
}

var BDialakoro迪亚拉科罗 feud.Barony = &迪亚拉科罗DialakoroBarony{}

func init() {
    f := BDialakoro迪亚拉科罗.(*迪亚拉科罗DialakoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dialakoro",
		TitleName: "迪亚拉科罗",
		TitleCode: "b_dialakoro",
	}
}
