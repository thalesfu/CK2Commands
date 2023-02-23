package calatayud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努埃瓦洛斯NuevalosBarony struct {
	feud.BaseBarony
}

var BNuevalos努埃瓦洛斯 feud.Barony = &努埃瓦洛斯NuevalosBarony{}

func init() {
	f := BNuevalos努埃瓦洛斯.(*努埃瓦洛斯NuevalosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nuevalos",
		TitleName: "努埃瓦洛斯",
		TitleCode: "b_nuevalos",
	}
}
