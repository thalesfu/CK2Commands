package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博迪亚姆BodiamBarony struct {
	feud.BaseBarony
}

var BBodiam博迪亚姆 feud.Barony = &博迪亚姆BodiamBarony{}

func init() {
	f := BBodiam博迪亚姆.(*博迪亚姆BodiamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodiam",
		TitleName: "博迪亚姆",
		TitleCode: "b_bodiam",
	}
}
