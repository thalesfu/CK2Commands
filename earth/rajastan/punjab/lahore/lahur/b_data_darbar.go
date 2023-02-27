package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达塔达巴尔Data_darbarBarony struct {
	feud.BaseBarony
}

var BData_darbar达塔达巴尔 feud.Barony = &达塔达巴尔Data_darbarBarony{}

func init() {
    f := BData_darbar达塔达巴尔.(*达塔达巴尔Data_darbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "data_darbar",
		TitleName: "达塔达巴尔",
		TitleCode: "b_data_darbar",
	}
}
