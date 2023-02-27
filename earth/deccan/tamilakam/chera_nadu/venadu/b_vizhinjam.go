package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗只旬VizhinjamBarony struct {
	feud.BaseBarony
}

var BVizhinjam毗只旬 feud.Barony = &毗只旬VizhinjamBarony{}

func init() {
    f := BVizhinjam毗只旬.(*毗只旬VizhinjamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vizhinjam",
		TitleName: "毗只旬",
		TitleCode: "b_vizhinjam",
	}
}
