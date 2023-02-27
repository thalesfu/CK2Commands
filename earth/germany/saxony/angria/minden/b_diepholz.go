package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪普霍尔茨DiepholzBarony struct {
	feud.BaseBarony
}

var BDiepholz迪普霍尔茨 feud.Barony = &迪普霍尔茨DiepholzBarony{}

func init() {
    f := BDiepholz迪普霍尔茨.(*迪普霍尔茨DiepholzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diepholz",
		TitleName: "迪普霍尔茨",
		TitleCode: "b_diepholz",
	}
}
