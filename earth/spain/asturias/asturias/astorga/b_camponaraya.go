package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎波那拉亚CamponarayaBarony struct {
	feud.BaseBarony
}

var BCamponaraya坎波那拉亚 feud.Barony = &坎波那拉亚CamponarayaBarony{}

func init() {
    f := BCamponaraya坎波那拉亚.(*坎波那拉亚CamponarayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "camponaraya",
		TitleName: "坎波那拉亚",
		TitleCode: "b_camponaraya",
	}
}
