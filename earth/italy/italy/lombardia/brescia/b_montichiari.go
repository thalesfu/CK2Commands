package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙蒂基亚里MontichiariBarony struct {
	feud.BaseBarony
}

var BMontichiari蒙蒂基亚里 feud.Barony = &蒙蒂基亚里MontichiariBarony{}

func init() {
    f := BMontichiari蒙蒂基亚里.(*蒙蒂基亚里MontichiariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montichiari",
		TitleName: "蒙蒂基亚里",
		TitleCode: "b_montichiari",
	}
}
