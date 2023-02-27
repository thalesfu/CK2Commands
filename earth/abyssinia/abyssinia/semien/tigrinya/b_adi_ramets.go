package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿迪拉梅兹Adi_rametsBarony struct {
	feud.BaseBarony
}

var BAdi_ramets阿迪拉梅兹 feud.Barony = &阿迪拉梅兹Adi_rametsBarony{}

func init() {
    f := BAdi_ramets阿迪拉梅兹.(*阿迪拉梅兹Adi_rametsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adi_ramets",
		TitleName: "阿迪拉梅兹",
		TitleCode: "b_adi_ramets",
	}
}
