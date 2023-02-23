package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎皮利亚CampigliaBarony struct {
	feud.BaseBarony
}

var BCampiglia坎皮利亚 feud.Barony = &坎皮利亚CampigliaBarony{}

func init() {
	f := BCampiglia坎皮利亚.(*坎皮利亚CampigliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "campiglia",
		TitleName: "坎皮利亚",
		TitleCode: "b_campiglia",
	}
}
