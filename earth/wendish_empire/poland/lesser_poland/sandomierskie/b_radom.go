package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉多姆RadomBarony struct {
	feud.BaseBarony
}

var BRadom拉多姆 feud.Barony = &拉多姆RadomBarony{}

func init() {
    f := BRadom拉多姆.(*拉多姆RadomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radom",
		TitleName: "拉多姆",
		TitleCode: "b_radom",
	}
}
