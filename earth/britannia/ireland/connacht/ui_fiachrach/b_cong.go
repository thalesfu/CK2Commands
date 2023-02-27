package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康镇CongBarony struct {
	feud.BaseBarony
}

var BCong康镇 feud.Barony = &康镇CongBarony{}

func init() {
    f := BCong康镇.(*康镇CongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cong",
		TitleName: "康镇",
		TitleCode: "b_cong",
	}
}
