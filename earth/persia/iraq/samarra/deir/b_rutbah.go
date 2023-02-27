package deir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁特巴赫RutbahBarony struct {
	feud.BaseBarony
}

var BRutbah鲁特巴赫 feud.Barony = &鲁特巴赫RutbahBarony{}

func init() {
    f := BRutbah鲁特巴赫.(*鲁特巴赫RutbahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rutbah",
		TitleName: "鲁特巴赫",
		TitleCode: "b_rutbah",
	}
}
