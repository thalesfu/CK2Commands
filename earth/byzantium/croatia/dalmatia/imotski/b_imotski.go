package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊莫茨基ImotskiBarony struct {
	feud.BaseBarony
}

var BImotski伊莫茨基 feud.Barony = &伊莫茨基ImotskiBarony{}

func init() {
	f := BImotski伊莫茨基.(*伊莫茨基ImotskiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "imotski",
		TitleName: "伊莫茨基",
		TitleCode: "b_imotski",
	}
}
