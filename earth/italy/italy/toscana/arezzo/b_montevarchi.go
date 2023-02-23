package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特瓦尔基MontevarchiBarony struct {
	feud.BaseBarony
}

var BMontevarchi蒙特瓦尔基 feud.Barony = &蒙特瓦尔基MontevarchiBarony{}

func init() {
	f := BMontevarchi蒙特瓦尔基.(*蒙特瓦尔基MontevarchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montevarchi",
		TitleName: "蒙特瓦尔基",
		TitleCode: "b_montevarchi",
	}
}
