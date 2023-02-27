package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙莱里MontlheryBarony struct {
	feud.BaseBarony
}

var BMontlhery蒙莱里 feud.Barony = &蒙莱里MontlheryBarony{}

func init() {
    f := BMontlhery蒙莱里.(*蒙莱里MontlheryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montlhery",
		TitleName: "蒙莱里",
		TitleCode: "b_montlhery",
	}
}
