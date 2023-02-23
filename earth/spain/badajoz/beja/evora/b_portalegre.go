package evora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波塔莱格雷PortalegreBarony struct {
	feud.BaseBarony
}

var BPortalegre波塔莱格雷 feud.Barony = &波塔莱格雷PortalegreBarony{}

func init() {
	f := BPortalegre波塔莱格雷.(*波塔莱格雷PortalegreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portalegre",
		TitleName: "波塔莱格雷",
		TitleCode: "b_portalegre",
	}
}
