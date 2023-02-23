package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科迪戈罗CodigoroBarony struct {
	feud.BaseBarony
}

var BCodigoro科迪戈罗 feud.Barony = &科迪戈罗CodigoroBarony{}

func init() {
	f := BCodigoro科迪戈罗.(*科迪戈罗CodigoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "codigoro",
		TitleName: "科迪戈罗",
		TitleCode: "b_codigoro",
	}
}
