package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法莱斯FalaiseBarony struct {
	feud.BaseBarony
}

var BFalaise法莱斯 feud.Barony = &法莱斯FalaiseBarony{}

func init() {
    f := BFalaise法莱斯.(*法莱斯FalaiseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "falaise",
		TitleName: "法莱斯",
		TitleCode: "b_falaise",
	}
}
