package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米达尔MidarBarony struct {
	feud.BaseBarony
}

var BMidar米达尔 feud.Barony = &米达尔MidarBarony{}

func init() {
    f := BMidar米达尔.(*米达尔MidarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "midar",
		TitleName: "米达尔",
		TitleCode: "b_midar",
	}
}
