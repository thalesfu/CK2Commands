package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特瑟内TeseneyBarony struct {
	feud.BaseBarony
}

var BTeseney特瑟内 feud.Barony = &特瑟内TeseneyBarony{}

func init() {
    f := BTeseney特瑟内.(*特瑟内TeseneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teseney",
		TitleName: "特瑟内",
		TitleCode: "b_teseney",
	}
}
