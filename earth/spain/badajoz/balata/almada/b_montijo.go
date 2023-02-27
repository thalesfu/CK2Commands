package almada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙蒂霍MontijoBarony struct {
	feud.BaseBarony
}

var BMontijo蒙蒂霍 feud.Barony = &蒙蒂霍MontijoBarony{}

func init() {
    f := BMontijo蒙蒂霍.(*蒙蒂霍MontijoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montijo",
		TitleName: "蒙蒂霍",
		TitleCode: "b_montijo",
	}
}
