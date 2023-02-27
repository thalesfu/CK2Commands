package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图兰巴TulambaBarony struct {
	feud.BaseBarony
}

var BTulamba图兰巴 feud.Barony = &图兰巴TulambaBarony{}

func init() {
    f := BTulamba图兰巴.(*图兰巴TulambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tulamba",
		TitleName: "图兰巴",
		TitleCode: "b_tulamba",
	}
}
