package mema

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨迪亚布古达SadiabougoudaBarony struct {
	feud.BaseBarony
}

var BSadiabougouda萨迪亚布古达 feud.Barony = &萨迪亚布古达SadiabougoudaBarony{}

func init() {
	f := BSadiabougouda萨迪亚布古达.(*萨迪亚布古达SadiabougoudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadiabougouda",
		TitleName: "萨迪亚布古达",
		TitleCode: "b_sadiabougouda",
	}
}
