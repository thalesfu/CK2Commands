package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南雪平SoderkopingBarony struct {
	feud.BaseBarony
}

var BSoderkoping南雪平 feud.Barony = &南雪平SoderkopingBarony{}

func init() {
    f := BSoderkoping南雪平.(*南雪平SoderkopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soderkoping",
		TitleName: "南雪平",
		TitleCode: "b_soderkoping",
	}
}
