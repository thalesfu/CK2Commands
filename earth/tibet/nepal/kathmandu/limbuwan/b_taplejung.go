package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达布莱宗TaplejungBarony struct {
	feud.BaseBarony
}

var BTaplejung达布莱宗 feud.Barony = &达布莱宗TaplejungBarony{}

func init() {
    f := BTaplejung达布莱宗.(*达布莱宗TaplejungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taplejung",
		TitleName: "达布莱宗",
		TitleCode: "b_taplejung",
	}
}
