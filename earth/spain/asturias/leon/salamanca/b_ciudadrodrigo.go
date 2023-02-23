package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗德里戈城CiudadrodrigoBarony struct {
	feud.BaseBarony
}

var BCiudadrodrigo罗德里戈城 feud.Barony = &罗德里戈城CiudadrodrigoBarony{}

func init() {
	f := BCiudadrodrigo罗德里戈城.(*罗德里戈城CiudadrodrigoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ciudadrodrigo",
		TitleName: "罗德里戈城",
		TitleCode: "b_ciudadrodrigo",
	}
}
