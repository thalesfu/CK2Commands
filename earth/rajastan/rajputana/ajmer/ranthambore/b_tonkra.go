package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通卡TonkraBarony struct {
	feud.BaseBarony
}

var BTonkra通卡 feud.Barony = &通卡TonkraBarony{}

func init() {
    f := BTonkra通卡.(*通卡TonkraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonkra",
		TitleName: "通卡",
		TitleCode: "b_tonkra",
	}
}
