package naupactus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利佐里基LidorikiBarony struct {
	feud.BaseBarony
}

var BLidoriki利佐里基 feud.Barony = &利佐里基LidorikiBarony{}

func init() {
	f := BLidoriki利佐里基.(*利佐里基LidorikiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lidoriki",
		TitleName: "利佐里基",
		TitleCode: "b_lidoriki",
	}
}
