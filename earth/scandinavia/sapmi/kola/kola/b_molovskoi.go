package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫洛夫斯科伊MolovskoiBarony struct {
	feud.BaseBarony
}

var BMolovskoi莫洛夫斯科伊 feud.Barony = &莫洛夫斯科伊MolovskoiBarony{}

func init() {
	f := BMolovskoi莫洛夫斯科伊.(*莫洛夫斯科伊MolovskoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "molovskoi",
		TitleName: "莫洛夫斯科伊",
		TitleCode: "b_molovskoi",
	}
}
