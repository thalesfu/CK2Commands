package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉古耶瓦茨KragujevacBarony struct {
	feud.BaseBarony
}

var BKragujevac克拉古耶瓦茨 feud.Barony = &克拉古耶瓦茨KragujevacBarony{}

func init() {
    f := BKragujevac克拉古耶瓦茨.(*克拉古耶瓦茨KragujevacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kragujevac",
		TitleName: "克拉古耶瓦茨",
		TitleCode: "b_kragujevac",
	}
}
