package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施塔莱克StahleckBarony struct {
	feud.BaseBarony
}

var BStahleck施塔莱克 feud.Barony = &施塔莱克StahleckBarony{}

func init() {
    f := BStahleck施塔莱克.(*施塔莱克StahleckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stahleck",
		TitleName: "施塔莱克",
		TitleCode: "b_stahleck",
	}
}
