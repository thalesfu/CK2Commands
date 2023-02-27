package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫里塞HriseyBarony struct {
	feud.BaseBarony
}

var BHrisey赫里塞 feud.Barony = &赫里塞HriseyBarony{}

func init() {
    f := BHrisey赫里塞.(*赫里塞HriseyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hrisey",
		TitleName: "赫里塞",
		TitleCode: "b_hrisey",
	}
}
